package config

import (
	"os"

	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`
	MongoDB         string `json:"mongo_db" mapstructure:"mongo_db"`
	ServiceDB       string `json:"service_db" mapstructure:"service_db"`
	JWTSecret       string `json:"jwt_secret" mapstructure:"jwt_secret"`
	DynamicDataGRPCAddr    string `json:"dynamic_data_grpc_addr" mapstructure:"dynamic_data_grpc_addr"`
	GenCodeGRPCAddr string `json:"gen_code_grpc_addr" mapstructure:"gen_code_grpc_addr"`
}

func loadDefaultConfig() *Config {

	return &Config{
		Base:        *config_lib.DefaultBaseConfig(),
		MongoDB:     "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB:   "user_service",
		JWTSecret:   os.Getenv("jwt_secret"),
		DynamicDataGRPCAddr:    "dynamic-data-service:443",
		GenCodeGRPCAddr: "gen-code-service:443",
	}
}
