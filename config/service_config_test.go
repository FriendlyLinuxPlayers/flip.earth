package config

import "testing"

type fakeConfig struct {
	FieldOne string `servconf:"field_one,required"`
	FieldTwo string `servconf:"field_two"`
}

func TestServiceConfig(t *testing.T) {
	var confValidSC1 ServiceConfig
	var confValidSC2 ServiceConfig
	var confInvalidSC1 ServiceConfig
	confValidSC1 = map[string]interface{}{
		"field_one": "test1",
		"field_two": "test2",
	}
	confValidSC2 = map[string]interface{}{
		"field_one": "test1",
	}
	confInvalidSC1 = map[string]interface{}{
		"field_two": "test",
	}

	// TODO add other config structs with invalid tags and test it
	t.Run("AssignValidSC1", func(t *testing.T) {
		cfg := fakeConfig{}
		err := confValidSC1.Assign(&cfg)
		if err != nil {
			t.Errorf("Assign returned an unexpected error: %s", err)
		}
		if cfg.FieldOne != "test1" {
			t.Errorf("FieldOne value %q is not \"test1\"", cfg.FieldOne)
		}
		if cfg.FieldTwo != "test2" {
			t.Errorf("FieldTwo value %q is not \"test2\"", cfg.FieldTwo)
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
		if cfg.FieldTwo != "" {
			t.Errorf("FieldTwo value %q is not \"\"", cfg.FieldTwo)
		}
	})
	t.Run("AssignInvalidSC1", func(t *testing.T) {
		cfg := fakeConfig{}
		err := confInvalidSC1.Assign(&cfg)
		if err == nil {
			t.Errorf("Assign did not return an error when missing a required field")
			return
		}
		switch err.(type) {
		case *InvalidTagError:
			if err.(*InvalidTagError).NameTag == "field_one" {
				return
			}
		}
		t.Errorf("Assign returned an unexpected error: %s", err)
	})
}
