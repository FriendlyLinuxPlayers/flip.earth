package service

import (
	"fmt"
	"strings"
)

// Builder is a simple implementation of ContainerBuilder.
type Builder struct {
	definitions []Definition
}

// Insert a new definition into the Builder.
func (b *Builder) Insert(def Definition) {
	b.definitions = append(b.definitions, def)
}

// Build creates the container once all definitions have been place in it. Note that
// dependencies must be inserted in order, for now.
func (b *Builder) Build() (Container, error) {
	if b.definitions == nil {
		return nil, fmt.Errorf("service: Definitions can not be nil")
	}
	numDefs := len(b.definitions)
	servs := make(map[string]interface{}, numDefs)
	for _, def := range b.definitions {
		if def.Dependencies == nil {
			def.Dependencies = make([]string, 0)
		}
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

		if def.Configuration == nil {
			def.Configuration = make(map[string]interface{}, 0)
		}

		if def.Init == nil {
			return nil, fmt.Errorf("service: Definitions must have an Initializer function.")
		}
		service, err := def.Init(deps, def.Configuration)
		if err != nil {
			return nil, err
		}

		def.Name = strings.TrimSpace(def.Name)
		if def.Name == "" {
			return nil, fmt.Errorf("service: service name must be non-empty and not whitespace only")
		}

		servs[def.Name] = service
	}

	return &SimpleContainer{
		services: servs,
	}, nil
}
