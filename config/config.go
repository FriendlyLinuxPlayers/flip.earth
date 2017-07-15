package config

// Config represents the basic application configuration.
type Config struct {
	Services map[string]ServiceConfig `json:"services"`
}
