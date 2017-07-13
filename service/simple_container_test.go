package service

import "testing"

func TestSimpleContainer(t *testing.T) {
	serviceMap := map[string]interface{}{
		"service": 1,
	}
	sc := SimpleContainer{
		services: serviceMap,
	}

	t.Run("GetValid", func(t *testing.T) {
		s, err := sc.Get("service")
		if err != nil {
			t.Errorf("Encountered an error: " + err.Error())
		}
		if s != serviceMap["service"] {
			t.Errorf("Returned service has the wrong value.")
		}
	})
	t.Run("GetInvalid", func(t *testing.T) {
		s, err := sc.Get("not a service")
		if err == nil {
			t.Errorf("Unexpectedly returned a service: %s", s)
		}
		switch err.(type) {
		case *MissingServiceError:
			t.SkipNow()
		}
		t.Errorf("Encountered an error: " + err.Error())
	})
	t.Run("HasValid", func(t *testing.T) {
		if !sc.Has("service") {
			t.Errorf("Does not have an expected service.")
		}
	})
	t.Run("HasInvalid", func(t *testing.T) {
		if sc.Has("not a service") {
			t.Errorf("Unexpectedly has a service.")
		}
	})
}
