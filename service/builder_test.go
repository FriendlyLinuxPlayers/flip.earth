package service

import "testing"

func fakeInit(deps, conf map[string]interface{}) (interface{}, error) {
	return "fakeInit", nil
}
func fakeDep(deps, conf map[string]interface{}) (interface{}, error) {
	return "fakeDep", nil
}

func TestBuilder(t *testing.T) {
	emptyDef := Definition{}
	dependencyDef := Definition{
		Name:          "dep service",
		Configuration: map[string]interface{}{"config": 1},
		Init:          fakeDep,
	}
	fullDef := Definition{
		Name:          "service",
		Dependencies:  []string{"dep service"},
		Configuration: map[string]interface{}{"config": 1},
		Init:          fakeInit,
	}
	// TODO add error checking function to Builder so that this isn't needed

	t.Run("Insert1", func(t *testing.T) {
		b := Builder{}
		b.Insert(emptyDef)
		if len(b.definitions) != 1 {
			t.Errorf("Incorrect number of definitions")
		}
	})
	t.Run("Insert2", func(t *testing.T) {
		b := Builder{}
		b.Insert(emptyDef)
		b.Insert(emptyDef)
		if len(b.definitions) != 2 {
			t.Errorf("Incorrect number of definitions")
		}
	})
	// TODO Test every possible combination of fields initialized in a Definition.
	t.Run("BuildInitDef1", func(t *testing.T) {
		b := Builder{
			definitions: []Definition{dependencyDef},
		}
		_, err := b.Build()
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildInitDef2", func(t *testing.T) {
		b := Builder{
			definitions: []Definition{dependencyDef, fullDef},
		}
		_, err := b.Build()
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildInitInvalidDef1", func(t *testing.T) {
		b := Builder{
			definitions: []Definition{fullDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when there were unmet dependencies (not present).")
		}
		switch err.(type) {
		case *MissingDepError:
			t.SkipNow()
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("BuildInitInvalidDef2", func(t *testing.T) {
		b := Builder{
			definitions: []Definition{fullDef, dependencyDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when there were unmet dependencies (wrong order).")
		}
		switch err.(type) {
		case *MissingDepError:
			t.SkipNow()
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("BuildEmpty", func(t *testing.T) {
		b := Builder{}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when no Definitions are present.")
		}
		if err != ErrNilDef {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
	t.Run("BuildUninitDef", func(t *testing.T) {
		b := Builder{
			definitions: []Definition{emptyDef},
		}
		_, err := b.Build()
		if err == nil {
			t.Errorf("Did not return an error when given an empty Definition.")
		}
		if err != ErrDefNilInit {
			t.Errorf("Encountered an error: " + err.Error())
		}
	})
}
