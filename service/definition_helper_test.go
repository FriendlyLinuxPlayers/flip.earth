package service

import "testing"

func TestDefinitionHelper(t *testing.T) {
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

	t.Run("InvalidReasonValid", func(t *testing.T) {
		t.Parallel()
		if err := InvalidReason(validDef); err != nil {
			t.Errorf("Valid Definition unexpectedly returned error: %s", err)
		}
	})
	t.Run("InvalidReasonVendor", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Vendor = ""
		if err := InvalidReason(def); err != ErrDefEmptyVendor {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("InvalidReasonPrefix", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Prefix = ""
		if err := InvalidReason(def); err != ErrDefEmptyPrefix {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("InvalidReasonName", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Name = ""
		if err := InvalidReason(def); err != ErrDefEmptyName {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("InvalidReasonInit", func(t *testing.T) {
		t.Parallel()
		def := validDef
		def.Init = nil
		if err := InvalidReason(def); err != ErrDefNilInit {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("InvalidReasonWhitespace", func(t *testing.T) {
		t.Parallel()
		err := InvalidReason(invalidWSDef)
		if err != ErrDefEmptyVendor && err != ErrDefEmptyPrefix && err != ErrDefEmptyName {
			t.Errorf("Definition returned an unexpected error: %s", err)
		}
	})
	t.Run("IsValidTrue1", func(t *testing.T) {
		t.Parallel()
		if !IsValid(validDef) {
			t.Errorf("Valid Definition unexpectedly returned false")
		}
	})
	t.Run("IsValidTrue2", func(t *testing.T) {
		t.Parallel()
		if !IsValid(validWSDef) {
			t.Errorf("Valid Definition unexpectedly returned false")
		}
	})
	t.Run("IsValidFalse", func(t *testing.T) {
		t.Parallel()
		if IsValid(invalidWSDef) {
			t.Errorf("Invalid Definition unexpectedly returned true")
		}
	})
	t.Run("TrimStrings1", func(t *testing.T) {
		t.Parallel()
		def := validDef
		TrimStrings(&def)
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
	t.Run("TrimStrings2", func(t *testing.T) {
		t.Parallel()
		def := validWSDef
		TrimStrings(&def)
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
	t.Run("TrimStrings3", func(t *testing.T) {
		t.Parallel()
		def := invalidWSDef
		TrimStrings(&def)
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
