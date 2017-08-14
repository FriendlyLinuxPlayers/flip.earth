package service

import (
	"reflect"
	"testing"
)

func TestSimpleContainer(t *testing.T) {
	typ := reflect.TypeOf(1)
	serviceByNameMap := map[string]interface{}{
		"service": 1,
	}
	serviceNameByTypeMap := map[reflect.Type]string{
		typ: "service",
	}
	sc := SimpleContainer{
		servicesByName:    serviceByNameMap,
		serviceNameByType: serviceNameByTypeMap,
	}

	t.Run("GetValid", func(t *testing.T) {
		t.Parallel()
		s, err := sc.Get("service")
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
		if s != serviceByNameMap["service"] {
			t.Errorf("Returned service has the wrong value: %+v", serviceByNameMap["service"])
		}
	})
	t.Run("GetInvalid", func(t *testing.T) {
		t.Parallel()
		s, err := sc.Get("not a service")
		if err == nil {
			t.Errorf("Unexpectedly returned a service: %+v", s)
		}
		switch err.(type) {
		case *MissingServiceError:
			return
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("HasValid", func(t *testing.T) {
		t.Parallel()
		if !sc.Has("service") {
			t.Errorf("Does not have an expected service: %q", "service")
		}
	})
	t.Run("HasInvalid", func(t *testing.T) {
		t.Parallel()
		if sc.Has("not a service") {
			t.Errorf("Unexpectedly has a service: %q", "not a service")
		}
	})
	t.Run("HasInvalidType", func(t *testing.T) {
		t.Parallel()
		if sc.HasAssignable(struct {
			secretName string
		}{
			"Anonymous Struct that can't possibly exist elsewhere",
		}) {
			t.Errorf("Unexpectedly has service for a type it shouldn't have")
		}
	})
	t.Run("HasValidType", func(t *testing.T) {
		t.Parallel()
		var i int
		if !sc.HasAssignable(i) {
			t.Errorf("Does not have an expected service for integer type")
		}
	})
	t.Run("GetValidType", func(t *testing.T) {
		t.Parallel()
		i := -99
		err := sc.Assign(&i)
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
		if i == 99 {
			t.Error("Failed to assign anything to passed reference")
		}
		if i != 1 {
			t.Errorf("Assigned incorrect value to passed reference")
		}
	})
	t.Run("GetInvalidType", func(t *testing.T) {
		t.Parallel()
		i := 99.99
		err := sc.Assign(&i)
		if err == nil {
			t.Errorf("failed to return an error for invalid type")
		}
	})
}
