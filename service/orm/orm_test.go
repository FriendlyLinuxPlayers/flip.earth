package orm

import "testing"

func TestOrm(t *testing.T) {
	testConfig := map[string]interface{}{
		"connection_string": "test_data/db/empty_test.db",
		"driver":            "sqlite3",
	}
	emptyDeps := make(map[string]interface{})

	t.Run("Init", func(t *testing.T) {
		_, err := Init(emptyDeps, testConfig)
		if err != nil {
			t.Errorf("Encountered an error: %s", err.Error())
		}
	})
}
