package service

import (
	"errors"
	"fmt"
)

var (
	ErrDefEmptyName   = errors.New("service: Definition name must be non-empty and not whitespace only")
	ErrDefEmptyPrefix = errors.New("service: Definition prefix must be non-empty and not witespace only")
	ErrDefEmptyVendor = errors.New("service: Definition vendor must be non-empty and not witespace only")
	ErrDefNilInit     = errors.New("service: Definition must have an Initializer function")
	ErrDefNilType     = errors.New("service: Definition must have a type")
	ErrNilDefs        = errors.New("service: \"definitions\" in ContainerBuilder cannot be nil")
)

const (
	missingDepPrefix     = "service: Could not find dependency %q for service %q"
	missingDepSuffix     = "Please make sure to insert them in order"
	missingServicePrefix = "service: Container does not have service"
)

// MissingDepError implements error, indicating that a dependency is missing
// from a ContainerBuilder.
type MissingDepError struct {
	// DepName is the name of the dependency which could not be found.
	DepName string
	// ServiceName is the name of the service which is missing the
	// dependency.
	ServiceName string
}

func (e *MissingDepError) Error() string {
	return fmt.Sprintf(missingDepPrefix+". %s", e.DepName, e.ServiceName, missingDepSuffix)
}

// MissingServiceError implements error, indicating that a service is missing
// from a service Container.
type MissingServiceError struct {
	// ServiceName is the name of the service which could not be found.
	ServiceName string
}

func (e *MissingServiceError) Error() string {
	return fmt.Sprintf("%s %q", missingServicePrefix, e.ServiceName)
}
