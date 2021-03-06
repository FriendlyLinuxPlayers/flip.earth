// FLiP.Earth is the website for the Friendly Linux Players community.
//
// See https://FriendlyLinuxPlayers.org for the live website and
// https://gitlab.com/FriendlyLinuxPlayers/flip.earth for the code.
package main

import (
	"fmt"

	_ "gitlab.com/friendlylinuxplayers/flip.earth/config"
	_ "gitlab.com/friendlylinuxplayers/flip.earth/router"
	_ "gitlab.com/friendlylinuxplayers/flip.earth/server"
	"gitlab.com/friendlylinuxplayers/flip.earth/service"
	cs "gitlab.com/friendlylinuxplayers/flip.earth/service/config"
)

// TODO refactor out everything so main only contains minimal code
func main() {
	b := new(service.Builder)
	configDef := service.Definition{
		Vendor: "flip",
		Prefix: "bootstrap",
		Name:   "config",
		Init:   cs.Init,
	}
	b.Insert(configDef)
	container, error := b.Build()
	if error != nil {
		panic(error)
	}

	service, error := container.Get("config")
	if error != nil {
		panic(error)
	}
	fmt.Printf("Config %+v \n", service)
}
