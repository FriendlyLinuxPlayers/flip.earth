package service

import (
	"testing"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

func fakeInit(deps, conf config.ServiceConfig) (interface{}, error) {
	return "fakeInit", nil
}
func fakeDep(deps, conf config.ServiceConfig) (interface{}, error) {
	return "fakeDep", nil
}

func TestBuilder(t *testing.T) {
	emptyDef := Definition{}
	dependencyDef := Definition{
		Name: "dep service",
		Init: fakeDep,
	}
	fullDef := Definition{
		Name:         "service",
		Dependencies: []string{"dep service"},
		Init:         fakeInit,
	}

	t.Run("Insert1", func(t *testing.T) {
		t.Parallel()
		b := Builder{}
		b.Insert(emptyDef)
		if len(b.definitions) != 1 {
			t.Errorf("Incorrect number of definitions. Expected: %d Result: %d", 1, len(b.definitions))
		}
	})
	t.Run("Insert2", func(t *testing.T) {
		t.Parallel()
		b := Builder{}
		b.Insert(emptyDef)
		b.Insert(emptyDef)
		if len(b.definitions) != 2 {
			t.Errorf("Incorrect number of definitions. Expected: %d Result: %d", 2, len(b.definitions))
		}
	})
	// TODO Test every possible combination of fields initialized in a Definition.
	t.Run("BuildInitDef1", func(t *testing.T) {
		t.Parallel()
		b := Builder{
			definitions: []Definition{dependencyDef},
		}
		_, err := b.Build()
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildInitDef2", func(t *testing.T) {
		t.Parallel()
		b := Builder{
			definitions: []Definition{dependencyDef, fullDef},
		}
		_, err := b.Build()
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildInitInvalidDef1", func(t *testing.T) {
		t.Parallel()
		b := Builder{
			definitions: []Definition{fullDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when there were unmet dependencies (not present).")
		}
		switch err.(type) {
		case *MissingDepError:
			return
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("BuildInitInvalidDef2", func(t *testing.T) {
		t.Parallel()
		b := Builder{
			definitions: []Definition{fullDef, dependencyDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when there were unmet dependencies (wrong order).")
		}
		switch err.(type) {
		case *MissingDepError:
			return
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("BuildEmpty", func(t *testing.T) {
		t.Parallel()
		b := Builder{}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when no Definitions are present.")
		}
		if err != ErrNilDefs {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildUninitDef", func(t *testing.T) {
		t.Parallel()
		b := Builder{
			definitions: []Definition{emptyDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when given an empty Definition.")
		}
		if err != ErrDefEmptyName && err != ErrDefNilInit {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
}
