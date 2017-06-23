package service

// Container is the basic interface which is to be implemented by anything providing
// a service container.
type Container interface {
	// Get should return the service which corresponds to the given name. If the
	// service doesn't exist, there should be an error returned.
	Get(name string) (interface{}, error)
	// Has should return true if a service with the given name exists, false
	// otherwise.
	Has(name string) bool
}

// Definition contains the metadata and Initializer required to use a service.
type Definition struct {
	// Name is what the service should be referred to in the Container.
	Name string
	// Dependencies contains the Names of the other services this service
	// dependends on.
	Dependencies []string
	// Configuration stores the service settings. It can have basically any
	// structure.
	Configuration map[string]interface{}
	// Initializer is what actually returns the service, using the Dependencies
	// and Configuration.
	Init Initializer
}

// Initializer is the function to be provided by anything that wants to return a
// service.
type Initializer func(deps map[string]interface{}, conf map[string]interface{}) (interface{}, error)

// ContainerBuilder builds the service container by having service Definitions
// inserted into it and the build method called.
type ContainerBuilder interface {
	// Insert inserts a service Definition.
	Insert(def Definition)
	// Build should be called when all services have been inserted into the
	// builder, which then returns a finished, usable Container or an error.
	Build() (*Container, error)
}
