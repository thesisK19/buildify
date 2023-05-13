package config

import (
	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`
	MongoDB         string `json:"mongo_db" mapstructure:"mongo_db"`
	ServiceDB       string `json:"service_db" mapstructure:"service_db"`
	UserGRPCAddr    string `json:"user_grpc_addr" mapstructure:"user_grpc_addr"`
	GenCodeGRPCAddr string `json:"gen_code_grpc_addr" mapstructure:"gen_code_grpc_addr"`
}

func loadDefaultConfig() *Config {

	return &Config{
		Base:            *config_lib.DefaultBaseConfig(),
		MongoDB:         "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB:       "dynamic_data_service",
		UserGRPCAddr:    "user-service:443",
		GenCodeGRPCAddr: "gen-code-service:443",
	}
}
