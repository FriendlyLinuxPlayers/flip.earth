package service

import "testing"

func TestDefinition(t *testing.T) {
	validDef := Definition{
		Vendor: "flip",
		Prefix: "test",
		Name:   "service",
		Init:   fakeInit,
	}
	validWSDef := Definition{
		Vendor: " flip",
		Prefix: "\ttest\n",
		Name:   " service ",
		Init:   fakeInit,
	}
	invalidWSDef := Definition{
		Vendor: "\n",
		Prefix: "\t",
		Name:   " ",
		Init:   fakeInit,
	}

	t.Run("invalidReasonValid", func(t *testing.T) {
		t.Parallel()
		if err := validDef.invalidReason(); err != nil {
			t.Errorf("Valid Definition unexpectedly returned error: %s", err)
		}
	})
	t.Run("invalidReasonVendor", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Vendor = ""
		if err := def.invalidReason(); err != ErrDefEmptyVendor {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("invalidReasonPrefix", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Prefix = ""
		if err := def.invalidReason(); err != ErrDefEmptyPrefix {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("invalidReasonName", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Name = ""
		if err := def.invalidReason(); err != ErrDefEmptyName {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("invalidReasonInit", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Init = nil
		if err := def.invalidReason(); err != ErrDefNilInit {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("invalidReasonWhitespace", func(t *testing.T) {
		t.Parallel()
		err := invalidWSDef.invalidReason()
		if err != ErrDefEmptyVendor && err != ErrDefEmptyPrefix && err != ErrDefEmptyName {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("isValidTrue1", func(t *testing.T) {
		t.Parallel()
		if !validDef.isValid() {
			t.Errorf("Valid Definition unexpectedly returned false")
		}
	})
	t.Run("isValidTrue2", func(t *testing.T) {
		t.Parallel()
		if !validWSDef.isValid() {
			t.Errorf("Valid Definition unexpectedly returned false")
		}
	})
	t.Run("isValidFalse", func(t *testing.T) {
		t.Parallel()
		if invalidWSDef.isValid() {
			t.Errorf("Invalid Definition unexpectedly returned true")
		}
	})
	t.Run("trimStrings1", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.trimStrings()
		if def.Vendor != "flip" {
			t.Errorf("Vendor does not match expected value \"flip\"")
		}
		if def.Prefix != "test" {
			t.Errorf("Prefix does not match expected value \"test\"")
		}
		if def.Name != "service" {
			t.Errorf("Name does not match expected value \"service\"")
		}
	})
	t.Run("trimStrings2", func(t *testing.T) {
		t.Parallel()
		def := validWSDef
		def.trimStrings()
		if def.Vendor != "flip" {
			t.Errorf("Vendor does not match expected value \"flip\"")
		}
		if def.Prefix != "test" {
			t.Errorf("Prefix does not match expected value \"test\"")
		}
		if def.Name != "service" {
			t.Errorf("Name does not match expected value \"service\"")
		}
	})
	t.Run("trimStrings3", func(t *testing.T) {
		t.Parallel()
		def := invalidWSDef
		def.trimStrings()
		if def.Vendor != "" {
			t.Errorf("Vendor does not match expected value \"\"")
		}
		if def.Prefix != "" {
			t.Errorf("Prefix does not match expected value \"\"")
		}
		if def.Name != "" {
			t.Errorf("Name does not match expected value \"\"")
		}
	})
}
