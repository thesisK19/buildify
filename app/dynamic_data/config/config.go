package config

import (
	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`
	MongoDB         string `json:"mongo_db" mapstructure:"mongo_db"`
	ServiceDB       string `json:"service_db" mapstructure:"service_db"`
	GenCodeHost     string `json:"gen_code_host" mapstructure:"gen_code_host"`
}

func loadDefaultConfig() *Config {

	return &Config{
		MongoDB:     "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB:   "dynamic_data_service",
		Base:        *config_lib.DefaultBaseConfig(),
		GenCodeHost: "localhost:9093",
	}
}
