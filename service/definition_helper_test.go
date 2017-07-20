package service

import (
	"fmt"
	"testing"
)

// testDefVal stores a value for testing Definition.
type testDefVal struct {
	// value is what the field in Definition should be.
	value interface{}
	// isValid indicates the boolean that the value, in isolation, is
	// expected to cause IsValid to return.
	isValid bool
}

func TestDefinitionHelper(t *testing.T) {
	// Contains all values to be tested
	// TODO try all possible combinations against IsValid
	valTable := map[string][]testDefVal{
		"Vendor":       []testDefVal{{" flip ", true}, {"flip", true}, {" ", false}, {"", false}},
		"Prefix":       []testDefVal{{" test ", true}, {"test", true}, {" ", false}, {"", false}},
		"Name":         []testDefVal{{" serv ", true}, {"serv", true}, {" ", false}, {"", false}},
		"Dependencies": []testDefVal{{[]string{"flip.test.none"}, true}, {nil, true}},
		"Init":         []testDefVal{{fakeInit, true}, {fakeDep, true}, {nil, false}},
	}
	// Contains indices to Vendor, Prefix, and Name slices in valTable.
	// This is used to test TrimStrings. The first value is the input and
	// the second is what the expected output is.
	trimTable := [][2]int{
		[2]int{0, 1},
		[2]int{2, 3},
		[2]int{1, 1},
		[2]int{3, 3},
	}
	sFlds := [3]string{"Vendor", "Prefix", "Name"}

	// Test trimTable against TrimStrings
	for i, set := range trimTable {
		vals := map[string][]string{
			sFlds[0]: make([]string, 2),
			sFlds[1]: make([]string, 2),
			sFlds[2]: make([]string, 2),
		}
		// Store strings from indices
		for _, f := range sFlds {
			// Input
			vals[f][0] = valTable[f][set[0]].value.(string)
			// Expected output
			vals[f][1] = valTable[f][set[1]].value.(string)
		}
		// Run the test
		t.Run(fmt.Sprintf("Trim%d", i+1), func(t *testing.T) {
			t.Parallel()
			def := Definition{vals[sFlds[0]][0], vals[sFlds[1]][0], vals[sFlds[2]][0], nil, nil}
			TrimStrings(&def)
			if def.Vendor != vals[sFlds[0]][1] {
				t.Errorf("%s %q does not match expected value %q", sFlds[0], def.Vendor, vals[sFlds[0]][1])
			}
			if def.Prefix != vals[sFlds[1]][1] {
				t.Errorf("%s %q does not match expected value %q", sFlds[1], def.Prefix, vals[sFlds[1]][1])
			}
			if def.Name != vals[sFlds[2]][1] {
				t.Errorf("%s %q does not match expected value %q", sFlds[2], def.Name, vals[sFlds[2]][1])
			}
		})
	}
}
