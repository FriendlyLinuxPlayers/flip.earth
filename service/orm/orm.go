package orm

import (
	"github.com/friendlylinuxplayers/flip.earth/config"
	"github.com/jinzhu/gorm"
	// required driver for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ormConfig struct {
	ConnectionString string `servconf:"connection_string,required"`
	Driver           string `servconf:"driver,required"`
}

//Init intializes and returns gorm
func Init(deps, conf config.ServiceConfig) (*gorm.DB, error) {
	cfg := ormConfig{}
	err := conf.Assign(&cfg)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(cfg.Driver, cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
