package main

import (
	"fmt"

	_ "github.com/friendlylinuxplayers/flip.earth/config"
	_ "github.com/friendlylinuxplayers/flip.earth/router"
	_ "github.com/friendlylinuxplayers/flip.earth/server"
	"github.com/friendlylinuxplayers/flip.earth/service"
	"github.com/friendlylinuxplayers/flip.earth/service/config"
)

// TODO refactor out everything so main only contains minimal code
func main() {
	b := new(service.Builder)
	configDef := service.Definition{"config", make([]string, 0), make(map[string]interface{}), cservice.Reader{}}
	b.Insert(configDef)
	container, error := b.Build()
	if error != nil {
		panic(error)
		return
	}

	service, error := container.Get("config")
	if error != nil {
		panic(error)
		return
	}
	fmt.Printf("Config %+v \n", service)
}
