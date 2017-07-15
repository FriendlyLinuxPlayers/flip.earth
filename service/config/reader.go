package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/friendlylinuxplayers/flip.earth/config"
)

// Init returns the Config service, a struct containing the parsed data.
//
// If a key "config_file" is present in conf, its value will be used as a custom
// path to the user-provided configuration file. A value of "config/config.json"
// indicates to search in the default location for a user config file. A value
// of "" (an empty string) indicates that only the default configuration should
// be used. If the value is invalid, an error is returned.
//
// User-provided configuration files will be merged with the default config file
// at "config/default/config.json" to set values which are missing.
//
// If a key "working_dir" is present in conf, its value will be used as a
// prefix to all relative paths used in locating configuration files. A value of
// "" indicates that relative paths will remain unchanged, otherwise it must be
// a valid directory.
func Init(deps, conf config.ServiceConfig) (interface{}, error) {
	configPath, ok := valToString(conf["config_file"], "config/config.json")
	if !ok {
		return nil, &ErrorPathInvalid{Cause: TypeAssert}
	}
	prefixDir, ok := valToString(conf["working_dir"], "")
	if !ok {
		return nil, &ErrorDirInvalid{Cause: TypeAssert}
	}
	if prefixDir != "" {
		dirInfo, err := os.Stat(prefixDir)
		if os.IsNotExist(err) {
			return nil, &ErrorDirInvalid{prefixDir, NotExist, err}
		}
		if !dirInfo.IsDir() {
			return nil, &ErrorDirInvalid{prefixDir, NotDir, nil}
		}
	}
	if configPath != "" && configPath != "config/config.json" {
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			return nil, &ErrorPathInvalid{configPath, NotExist, err}
		}
		// If configPath is relative, append it to prefixDir
		if !os.IsPathSeparator(configPath[0]) {
			configPath = prefixDir + configPath
		}
	}
	return parseConfig(configPath, prefixDir)
}

// valToString takes an interface and converts it to a string. If value is nil,
// the provided def is returned.
func valToString(value interface{}, def string) (string, bool) {
	if value == nil {
		return def, true
	}
	sVal, ok := value.(string)
	return sVal, ok
}

// parseConfig reads the default and user json config files, returning a Config.
func parseConfig(userConfig, relPrefix string) (*config.Config, error) {
	cPath := relPrefix + "config/default/config.json"
	cfgFile, err := ioutil.ReadFile(cPath)
	if err != nil {
		return nil, &ErrorParse{cPath, true, true, err}
	}
	cfg := new(config.Config)
	err = json.Unmarshal(cfgFile, cfg)
	if err != nil {
		return nil, &ErrorParse{cPath, true, false, err}
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
		return &ErrorParse{userConfigPath, false, true, err}
	}

	err = json.Unmarshal(userCfgFile, cfg)
	if err != nil {
		return &ErrorParse{userConfigPath, false, false, err}
	}

	return nil
}
