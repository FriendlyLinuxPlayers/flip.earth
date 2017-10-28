package service

import (
	"reflect"

	"github.com/friendlylinuxplayers/flip.earth/service/config"
)

type Definition interface {
	Type() reflect.Type
	SetConfig(interface{}) //TODO figure out proper type for service configuration
	Priority() int
}

type SingletonDefinition interface {
	Definition
	BuildSingleton() interface{} //TODO figure out proper type
}

type TransientDefinition interface {
	Definition
	BuildTransient() interface{} //TODO figure out proper type
}
