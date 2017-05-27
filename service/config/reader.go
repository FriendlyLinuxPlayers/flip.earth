package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

//Reader is the service for getting the application configuration. It cannot be
// rebuilt or changed after it is first used.
type Reader struct{}

//Init returns the config service, which is just a struct containing the data.
func (r Reader) Init(deps map[string]interface{}, conf map[string]interface{}) (interface{}, error) {
	return parseConfig()
}

//Reads the default and user json config files, returning a Config.
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

//Merges the user config.json data with the provided Config.
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
