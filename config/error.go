package config

import (
	"fmt"
	"reflect"
)

const (
	invalidArgumentPrefix = "config: Assign argument Kind %q is not expected Kind %s"
	emptyNameTagPrefix    = "config: \"servconf\" tag for the struct field %q is invalid: name field must be non-empty and not whitespace only"
	missingRequiredPrefix = "config: ServiceConfig is missing a required value for %q"
)

// InvalidArgumentError implements error, indicating that Assign's argument is
// not of the expected Kind.
type InvalidArgumentError struct {
	// Kind is the unexpected Kind.
	Kind reflect.Kind
	// Indirected indicates if the value has been indirected or not
	Indirected bool
}

func (e *InvalidArgumentError) Error() string {
	suffix := fmt.Sprintf("%q", reflect.Ptr)
	if e.Indirected {
		suffix = fmt.Sprintf("%q after being indirected", reflect.Struct)
	}
	return fmt.Sprintf(invalidArgumentPrefix, e.Kind, suffix)
}

// EmptyNameTagError implements error, indicating that the first tag field,
// which represents the tag name, is empty.
type EmptyNameTagError struct {
	// FieldName is the struct field that has an empty name.
	FieldName string
}

func (e *EmptyNameTagError) Error() string {
	return fmt.Sprintf(emptyNameTagPrefix, e.FieldName)
}

// MissingRequiredError implements error, indicating that the ServiceConfig
// lacks a key:value pair corresponding to a required field.
type MissingRequiredError struct {
	// NameTag is the name tag/key which does not have a required value.
	NameTag string
}

func (e *MissingRequiredError) Error() string {
	return fmt.Sprintf(missingRequiredPrefix, e.NameTag)
}
