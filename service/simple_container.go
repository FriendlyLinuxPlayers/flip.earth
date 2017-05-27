package service

import "fmt"

type SimpleContainer struct {
	services map[string]interface{}
}

func (sc *SimpleContainer) Get(name string) (interface{}, error) {
	if s, ok := sc.services[name]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("service: Container does not have service %q", name)
}

func (sc *SimpleContainer) Has(name string) bool {
	_, ok := sc.services[name]
	return ok
}
