package config

import (
	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`

	MongoDB   string `json:"mongo_db" mapstructure:"mongo_db"`
	ServiceDB string `json:"service_db" mapstructure:"service_db"`

	UserGRPCAddr        string `json:"user_grpc_addr" mapstructure:"user_grpc_addr"`
	DynamicDataGRPCAddr string `json:"dynamic_data_grpc_addr" mapstructure:"dynamic_data_grpc_addr"`
}

func loadDefaultConfig() *Config {
	return &Config{
		Base:                *config_lib.DefaultBaseConfig(),
		MongoDB:             "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB:           "gen_code_service",
		UserGRPCAddr:        "user-service:443",
		DynamicDataGRPCAddr: "dynamic-data-service:443",
	}
}
