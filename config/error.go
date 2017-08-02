package config

import (
	"fmt"
	"reflect"
)

const (
	invalidArgumentPrefix = "config: Assign argument Kind %q is not expected Kind %s"
	invalidTagPrefix      = "config: \"servconf\" tag for the struct field %q is invalid: %s"
)

// InvalidArgumentError implements error, indicating that the value of "to" is not of the expected Kind.
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

// InvalidTagError implements error, indicating that there is a problem with the
// field tags in the struct provided to ServiceConfig.Assign.
type InvalidTagError struct {
	// FieldName is the name of the field in the struct.
	FieldName string
	// NameTag is the value of the first tag field. If empty, the error
	// indicates that the name field must not be empty. If not empty, the
	// error indicates that the tag is required but missing.
	NameTag string
}

func (e *InvalidTagError) Error() string {
	if e.NameTag == "" {
		return fmt.Sprintf(invalidTagPrefix, e.FieldName, "name field must be non-empty and not whitespace only")
	}
	return fmt.Sprintf(invalidTagPrefix+"%q", e.FieldName, "missing required field ", e.NameTag)
}
