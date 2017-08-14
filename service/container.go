package service

// Container is the basic interface which is to be implemented by anything
// providing a service container.
type Container interface {
	// Get should return the service which corresponds to the given name. If
	// the service doesn't exist, there should be an error returned.
	Get(name string) (interface{}, error)
	// Has should return true if a service with the given name exists, false
	// otherwise.
	Has(name string) bool
	// Assign should take a reference to a variable and should assign the service,
	// which corresponds to that variable's type. If such a service doesn't exist
	// an error is returned
	Assign(to interface{}) error
	// Has Assignable should check if the type of the passed variable can be assigned to
	// and return the corresponding bool value
	HasAssignable(typ interface{}) bool
}

// ContainerBuilder builds the service Container by having service Definitions
// inserted into it and the build method called.
type ContainerBuilder interface {
	// Insert inserts a service Definition.
	Insert(def Definition)
	// Build should be called when all services have been inserted into the
	// builder, which then returns a finished, usable Container or an error.
	Build() (*Container, error)
}
