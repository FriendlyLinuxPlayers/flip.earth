package service

import "fmt"

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
	return nil, fmt.Errorf("service: SimpleContainer does not have service %q", name)
}

// Has checks if a service under the given name exists within the container.
func (sc *SimpleContainer) Has(name string) bool {
	_, ok := sc.services[name]
	return ok
}
