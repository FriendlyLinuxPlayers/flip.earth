package service

import "github.com/friendlylinuxplayers/flip.earth/config"

// Builder is a simple implementation of ContainerBuilder.
type Builder struct {
	ServiceConfigs map[string]config.ServiceConfig
	definitions    []Definition
}

// Insert a new definition into the Builder.
func (b *Builder) Insert(def Definition) error {
	if !IsValid(def) {
		return InvalidReason(def)
	}
	b.definitions = append(b.definitions, def)
	return nil
}

// Build creates the container once all definitions have been place in it. Note
// that dependencies must be inserted in order, for now.
func (b *Builder) Build() (Container, error) {
	if b.definitions == nil {
		return nil, ErrNilDefs
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
				return nil, &MissingDepError{name, def.Name}
			}
			deps[name] = dep
		}
		conf, ok := b.ServiceConfigs[def.Vendor+"."+def.Prefix+"."+def.Name]
		if !ok {
			conf = config.ServiceConfig{}
		}

		service, err := def.Init(deps, conf)
		if err != nil {
			return nil, err
		}

		servs[def.Name] = service
	}

	return &SimpleContainer{
		services: servs,
	}, nil
}
