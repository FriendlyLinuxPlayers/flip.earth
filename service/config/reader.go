package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

// Init returns the Config service, a struct containing the parsed data.
//
// If a key "config_file" is present in conf, its value will be used as a custom
// path to the user-provided configuration file. A value of "" (an empty string)
// indicates that only the default configuration should be used. If the value is
// invalid, an error is returned. Otherwise, the user config path defaults to
// "config/config.json"
func Init(deps, conf map[string]interface{}) (interface{}, error) {
	configVal, ok := conf["config_file"]
	if !ok {
		return parseConfig("config/config.json")
	}
	configPath, ok := configVal.(string)
	if !ok {
		return nil, fmt.Errorf("user provided config file path is invalid: string type assertion failed")
	}
	if configPath == "" {
		return parseConfig("")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("user provided config file does not exist: %s", configPath)
	}
	return parseConfig(configPath)
}

// parseConfig reads the default and user json config files, returning a Config.
func parseConfig(userConfig string) (*config.Config, error) {
	cfgFile, err := ioutil.ReadFile("config/default/config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading default config file: %s", err.Error())
	}
	cfg := new(config.Config)
	err = json.Unmarshal(cfgFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing default config file: %s", err.Error())
	}

	err = mergeUserConfig(cfg, userConfig)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// mergeUserConfig adds the user json config data to the provided Config, if it
// exists.
func mergeUserConfig(cfg *config.Config, userConfigPath string) error {
	if _, err := os.Stat(userConfigPath); os.IsNotExist(err) {
		return nil
	}

	userCfgFile, err := ioutil.ReadFile(userConfigPath)
	if err != nil {
		return fmt.Errorf("error reading user config file: %s", err.Error())
	}

	err = json.Unmarshal(userCfgFile, cfg)
	if err != nil {
		return fmt.Errorf("error parsing user config file: %s", err.Error())
	}

	return nil
}
