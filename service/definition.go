package service

import (
	"strings"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

// Initializer is the function to be provided by anything that wants to return a
// service.
type Initializer func(deps, conf config.ServiceConfig) (interface{}, error)

// Definition contains the metadata and Initializer required to use a service.
// All valid Definitions require the Name and Init fields, the rest are
// optional.
type Definition struct {
	// Vendor is the vendor prefix for the service name.
	// Leading and trailing whitespace will be trimmed.
	// This is the first step torwards service decoration.
	Vendor string
	// Prefix.
	// Leading and trailing whitespace will be trimmed.
	Prefix string
	// Name is what the service should be referred to in the Container.
	// Leading and trailing whitespace will be trimmed.
	Name string
	// Dependencies contains the Names of the other services this service
	// dependends on.
	Dependencies []string
	// Init is what actually returns the service, using the Dependencies and
	// Configuration.
	Init Initializer
}

// trimStrings trims the whitespace from Vendor, Prefix, and Name fields in the
// Definition.
func (d *Definition) trimStrings() {
	d.Name = strings.TrimSpace(d.Name)
	d.Prefix = strings.TrimSpace(d.Prefix)
	d.Vendor = strings.TrimSpace(d.Vendor)
}

// isValid returns a boolean indicating if the Definition's fields meet the
// minimum requirements.
func (d *Definition) isValid() bool {
	if d.invalidReason() == nil {
		return true
	}
	return false
}

// invalidReason returns an error explaining what makes the Definition invalid,
// or nil if the Definition is valid.
func (d *Definition) invalidReason() error {
	name := strings.TrimSpace(d.Name)
	if name == "" {
		return ErrDefEmptyName
	}
	prefix := strings.TrimSpace(d.Prefix)
	if prefix == "" {
		return ErrDefEmptyPrefix
	}
	vendor := strings.TrimSpace(d.Vendor)
	if vendor == "" {
		return ErrDefEmptyVendor
	}
	if d.Init == nil {
		return ErrDefNilInit
	}
	return nil
}
