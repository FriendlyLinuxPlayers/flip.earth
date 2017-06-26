package service

import (
	"fmt"
	"strings"
)

// SimpleContainer is a simple implementation of Container.
type SimpleContainer struct {
	services map[string]interface{}
}

// Get returns the service with the matching name. If it doesn't exist, an error is
// returned.
func (sc *SimpleContainer) Get(name string) (interface{}, error) {
	if s, ok := sc.services[name]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("%s %q", notPresentErrorPrefix(), name)
}

// Has checks if a service under the given name exists within the container.
func (sc *SimpleContainer) Has(name string) bool {
	_, ok := sc.services[name]
	return ok
}

// IsNotPresent returns a boolean indicating whether the error is known to report
// that a service is not contained in a SimpleContainer.
func IsNotPresent(err error) bool {
	return strings.Contains(err.Error(), notPresentErrorPrefix())
}

// notPresentErrorPrefix returns the string that prefixes the error returned when
// a service is not contained in a SimpleContainer.
func notPresentErrorPrefix() string {
	return "service: SimpleContainer does not have service"
}
