package config

// Config represents the basic application configuration.
type Config struct {
	Db DBConfig `json:"db"`
}

// DBConfig
type DBConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	DbName   string `json:"db_name"`
	Password string `json:"password"`
}
