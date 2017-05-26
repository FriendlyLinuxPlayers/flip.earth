package service

//Container is the basic interface which is to be implemented by anything providing a service container
type Container interface {
	//Get should return the service and which corresponds to the given name
	//if the service doesn't exist, there should be an error returned
	Get(name string) (interface{}, error)
	//Has should return true if a service with the given name exists, false otherwise
	Has(name string) bool
}

//Definition for a sevice
type Definition struct {
	//The Name of the service under which it should be stored in the container
	Name string
	//the names of the other services this service dependends on
	Dependencies []string
	//the configuration for the service, can have basically any structure
	Configuration map[string]interface{}
	// The Initializer of the service, this is what actually return the service
	Initializer Initializer
}

//Initializer interface to be implemented by anything that wants to return a service
type Initializer interface {
	//The method should be passed a string-indexed (the strings being the service names) map of fully working services
	// and the configuration as found in the Definition struct of the service. If the service can be successfully initialized
	// it should be returned, in case of an error during initialization an error is returned
	Init(deps map[string]interface{}, conf map[string]interface{}) (interface{}, error)
}

//ContainerBuilder builds the service container by having service Definitions inserted into
// and the build methid called
type ContainerBuilder interface {
	//Insert inserts a service Definition
	Insert(def Definition)
	//Build should be called when all services have been inserted into the builder
	//ith then returns a finished, usable Container or an error
	Build() (Container, error)
}
