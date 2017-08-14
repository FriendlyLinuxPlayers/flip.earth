package service

import (
	"fmt"
	"reflect"
)

// SimpleContainer is a simple implementation of Container.
type SimpleContainer struct {
	servicesByName    map[string]interface{}
	serviceNameByType map[reflect.Type]string
}

// Get returns the service with the matching name. If it doesn't exist, an error
// is returned.
func (sc *SimpleContainer) Get(name string) (interface{}, error) {
	if s, ok := sc.servicesByName[name]; ok {
		return s, nil
	}
	return nil, &MissingServiceError{name}
}

// Has checks if a service under the given name exists within the container.
func (sc *SimpleContainer) Has(name string) bool {
	_, ok := sc.servicesByName[name]
	return ok
}

// Assign attempts to assign the service to the passed argument, that corresponds with its underlying type
// if that isn't possible, an error is returned and the passed variable remains unassigned
func (sc *SimpleContainer) Assign(to interface{}) error {

	vTop := reflect.ValueOf(to)
	if vTop.Kind() != reflect.Ptr {
		return fmt.Errorf("'to' is not a pointer") //TODO Error types
	}
	t := reflect.TypeOf(to).Elem()
	name, err := sc.getNameForType(t)
	if err != nil {
		return err
	}

	service, err := sc.Get(name)
	if err != nil {
		return err
	}

	// This should hopefully always be a nonprobelamtic action because of how we build the container
	v := vTop.Elem()
	serviceV := reflect.ValueOf(service)
	v.Set(serviceV)

	return nil
}

func (sc *SimpleContainer) getNameForType(typ reflect.Type) (string, error) {
	name, hasType := sc.serviceNameByType[typ]
	if !hasType {
		return "", fmt.Errorf("simple container: No service name found for type %q", typ) //TODO error types
	}

	return name, nil
}

// HasAssignable checks if the container has a service for the type of the passed argument
func (sc *SimpleContainer) HasAssignable(typ interface{}) bool {
	t := reflect.TypeOf(typ)
	name, err := sc.getNameForType(t)
	if err != nil {
		return false
	}

	return sc.Has(name)

}
