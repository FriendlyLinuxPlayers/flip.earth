package config

import "testing"

type fakeConfig struct {
	FieldOne string `servconf:"field_one,required"`
	FieldTwo bool   `servconf:"field_two"`
	// unexported should be ignored by Assign, without processing tags.
	unexported string `servconf:"unexported"`
}

type fakeInvalidConfig struct {
	NoName string `servconf:" "`
}

func TestServiceConfig(t *testing.T) {
	var confValidSC1 ServiceConfig
	var confValidSC2 ServiceConfig
	var confInvalidSC ServiceConfig
	confValidSC1 = map[string]interface{}{
		"field_one":  "test1",
		"field_two":  true,
		"unexported": "not used",
	}
	confValidSC2 = map[string]interface{}{
		"field_one": "test1",
	}
	confInvalidSC = map[string]interface{}{
		"field_two": true,
	}

	// TODO test non-convertible values
	t.Run("AssignValidSC1", func(t *testing.T) {
		cfg := fakeConfig{}
		err := confValidSC1.Assign(&cfg)
		if err != nil {
			t.Errorf("Assign returned an unexpected error: %s", err)
		}
		if cfg.FieldOne != "test1" {
			t.Errorf("FieldOne value %q is not \"test1\"", cfg.FieldOne)
		}
		if !cfg.FieldTwo {
			t.Errorf("FieldTwo value %v is not true", cfg.FieldTwo)
		}
		if cfg.unexported != "" {
			t.Errorf("unexported value %q has been changed from \"\"", cfg.unexported)
		}
	})
	t.Run("AssignValidSC2", func(t *testing.T) {
		cfg := fakeConfig{}
		err := confValidSC2.Assign(&cfg)
		if err != nil {
			t.Errorf("Assign returned an unexpected error: %s", err)
		}
		if cfg.FieldOne != "test1" {
			t.Errorf("FieldOne value %q is not \"test1\"", cfg.FieldOne)
		}
		if cfg.FieldTwo {
			t.Errorf("FieldTwo value %v is not false", cfg.FieldTwo)
		}
		if cfg.unexported != "" {
			t.Errorf("unexported value %q has been changed from \"\"", cfg.unexported)
		}
	})
	t.Run("AssignInvalidSC", func(t *testing.T) {
		cfg := fakeConfig{}
		err := confInvalidSC.Assign(&cfg)
		if err == nil {
			t.Errorf("Assign did not return an error when missing a required field")
			return
		}
		switch err.(type) {
		case *MissingRequiredError:
			return
		}
		t.Errorf("Assign returned an unexpected error: %s", err)
	})
	t.Run("AssignInvalidStruct", func(t *testing.T) {
		cfg := fakeInvalidConfig{}
		err := confValidSC1.Assign(&cfg)
		if err == nil {
			t.Errorf("Assign did not return an error when missing a name tag")
			return
		}

	})
}
