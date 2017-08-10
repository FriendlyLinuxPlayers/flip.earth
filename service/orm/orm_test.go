package orm

import "testing"
import "github.com/mattn/go-sqlite3"

func TestOrm(t *testing.T) {
	testConfig := map[string]interface{}{
		"connection_string": "test_data/db/empty_test.db",
		"driver":            "sqlite3",
	}
	emptyDeps := make(map[string]interface{})

	// TODO Test a valid config (after setting up test_data/ as we want)
	t.Run("InitInvalidConf", func(t *testing.T) {
		conf := testConfig
		conf["connection_string"] = "test_data/db/not_exist.db"
		_, err := Init(emptyDeps, conf)
		if err == nil {
			t.Errorf("Invalid conf unexpectedly did not return an error.")
		}
		switch err.(type) {
		case sqlite3.Error:
			return
		}
		t.Errorf("Unexpected error: " + err.Error())
	})
}
