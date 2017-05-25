package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/friendlylinuxplayers/flip.earth/config"
	_ "github.com/friendlylinuxplayers/flip.earth/router"
	_ "github.com/friendlylinuxplayers/flip.earth/server"
)

// TODO refactor out everything so main only contains main func with minimal code
func main() {
	cfg, err := parseConfig()

	if err != nil {
		panic(err)
	}

	cfgVal := *cfg

	fmt.Printf("Config %+v \n", cfgVal)
}

func parseConfig() (*config.Config, error) {
	cfgFile, err := ioutil.ReadFile("config/default/config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading default config file: %s", err.Error())
	}
	cfg := new(config.Config)
	err = json.Unmarshal(cfgFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing default config file: %s", err.Error())
	}

	err = mergeUserConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func mergeUserConfig(cfg *config.Config) error {
	if _, err := os.Stat("config/config.json"); os.IsNotExist(err) {
		return nil
	}

	userCfgFile, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return fmt.Errorf("error reading user config file: %s", err.Error())
	}

	err = json.Unmarshal(userCfgFile, cfg)
	if err != nil {
		return fmt.Errorf("error parsing user config file: %s", err.Error())
	}

	return nil
}
