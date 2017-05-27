package service

import fmt

type SimpleContainer struct {
	services map[string]interface{}
}

func Get(name String) interface{}, error {
	if s, ok := services[name]; ok {
		return s
	}
	return _, fmt.Errorf("service: Container does not have service %q", name)
}

func Has(name String) bool {
	_, ok := services[name]
	return ok
}
