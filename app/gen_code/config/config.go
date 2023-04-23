package config

import (
	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`

	MongoDB   string `mapstructure:"mongo_db"`
	ServiceDB string `mapstructure:"service_db"`
}

func loadDefaultConfig() *Config {
	return &Config{
		MongoDB:   "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB: "gen_code_service",
		Base:      *config_lib.DefaultBaseConfig(),
	}
}
