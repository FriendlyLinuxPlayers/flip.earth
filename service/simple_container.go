package service

import (
	"fmt"
	"reflect"
)

type item interface {
	service() (interface{}, error)
}

type singletonItem struct {
	instance interface{}
}

func (si singletonItem) service() (interface{}, error) {
	return si.instance, nil
}

type transientItem struct {
	definition TransientDefinition
}

func (ti transientItem) service() (interface{}, error) {
	return ti.definition.BuildTransient()
}

// SimpleContainer is a simple implementation of Container.
type SimpleContainer struct {
	servicesByType map[reflect.Type]item
}

// Has checks if a service under the given name exists within the container.
func (sc *SimpleContainer) Has(typ reflect.Type) bool {
	_, ok := sc.servicesByType[typ]
	return ok
}

func (sc *SimpleContainer) Get(typ reflect.Type) (interface{}, error) {
	if hasType := sc.Has(typ); !hasType {
		return nil, fmt.Errorf("simple container: Container does not have type '%s'", typ) //TODO proper error type
	}

	service, err := sc.servicesByType[typ].service()
	if err != nil {
		return nil, err
	}

	return service, nil

}

// Assign attempts to assign the service to the passed argument, that corresponds with its underlying type
// if that isn't possible, an error is returned and the passed variable remains unassigned
func (sc *SimpleContainer) Assign(to interface{}) error {

	toValue := reflect.ValueOf(to)
	if toValue.Kind() != reflect.Ptr {
		return fmt.Errorf("'to' is not a pointer") //TODO Error types
	}
	underlyingType := reflect.TypeOf(to).Elem()

	service, err := sc.Get(underlyingType)

	if err != nil {
		return err
	}

	reflect.ValueOf(underlyingType).Set(reflect.ValueOf(service)) //TODO check if this is correct

	return nil
}

func (sc *SimpleContainer) Inject(structPtr interface{}) error {
	sPtrVal := reflect.ValueOf(structPtr)

	if sPtrVal.Kind() != reflect.Ptr {
		return fmt.Errorf("'structPtr' is not a a pointer") //TODO Error types
	}

	underlyingType := reflect.TypeOf(structPtr).Elem()
	numFields := underlyingType.NumField()
	underlyingValue := reflect.ValueOf(structPtr).Elem()
	for i := 0; i < numFields; i++ {
		field := underlyingType.Field(i)
		_, needsInject := field.Tag.Lookup("inject")
		if !needsInject {
			continue
		}

		valField := underlyingValue.Field(i)
		if !valField.CanSet() {
			return fmt.Errorf("simple container: Can't set field '%s' at index '%d' of type '%s'", field.Name, i, underlyingType)
		}

		service, err := sc.Get(field.Type)
		if err != nil {
			return err
		}

		valField.Set(reflect.ValueOf(service))
	}

	return nil
}
