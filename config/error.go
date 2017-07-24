package config

import (
	"errors"
	"fmt"
)

var ErrNotStruct = errors.New("config: Assign \"to\" interface must be a struct")

const invalidTagPrefix = "config: \"servconf\" tag for the struct field %q is invalid: %s"

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
