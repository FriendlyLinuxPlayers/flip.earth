package service

import "fmt"

//Builder is a simple implementation of ContainerBuilder
type Builder struct {
	definitions []Definition
}

//Insert a new definition into the Builder
func (b *Builder) Insert(def Definition) {
	b.definitions = append(b.definitions, def)
}

//Build builds the container once all definitions have been place in it
//dependencies need to be inserted in order for now
func (b *Builder) Build() (Container, error) {
	numDefs := len(b.definitions)
	servs := make(map[string]interface{}, numDefs)
	for _, def := range b.definitions {
		numDeps := len(def.Dependencies)
		deps := make(map[string]interface{}, numDeps)
		for _, name := range def.Dependencies {
			dep, ok := servs[name]
			if !ok {
				return nil, fmt.Errorf("service: Could not find "+
					"dependency %q for service %q. Please make "+
					"sure to insert them in order", name, def.Name)
			}
			deps[name] = dep
		}
		service, err := def.Initializer.Init(deps, def.Configuration)
		if err != nil {
			return nil, err
		}
		servs[def.Name] = service
	}

	return SimpleContainer{
		services: servs,
	}, nil
}
