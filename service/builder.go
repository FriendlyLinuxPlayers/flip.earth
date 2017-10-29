package service

//TODO refactor
import (
	"fmt"
	"reflect"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

// SimpleBuilder is a simple implementation of ContainerBuilder.
type SimpleBuilder struct {
	ServiceConfigs map[string]config.ServiceConfig //TODO config type is probably bad and needs to be changed
	definitions    []Definition
}

// Insert a new definition into the Builder.
func (b *SimpleBuilder) Insert(def ...Definition) error {
	for _, definition := range def {
		if invalidReason := isValid(definition); invalidReason != nil {
			return fmt.Errorf("container builder invalid definition: %s", invalidReason)
		}
	}
	b.definitions = append(b.definitions, def...)
	return nil
}

// Build creates the container based on the currently inserted definitions
func (b *SimpleBuilder) Build() (Container, error) {
	//TODO complete
	partialContainer := SimpleContainer{
		servicesByType: make(map[reflect.Type]item), //TODO remember changing the type here
	}
	//orderedDefinitions := make([]Definition, 0, len(b.definitions))

	return &partialContainer, nil
}

func isValid(def Definition) error {
	//TODO implement
	return nil
}
