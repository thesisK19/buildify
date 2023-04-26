package config

import (
	"os"

	config_lib "github.com/thesisK19/buildify/library/config"
)

type Config struct {
	config_lib.Base `mapstructure:",squash"`
	MongoDB         string `json:"mongo_db" mapstructure:"mongo_db"`
	ServiceDB       string `json:"service_db" mapstructure:"service_db"`
	GenCodeHost     string `json:"gen_code_host" mapstructure:"gen_code_host"`
	JWTSecret       string `json:"jwt_secret" mapstructure:"jwt_secret"`
}

func loadDefaultConfig() *Config {

	return &Config{
		MongoDB:     "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB:   "user_service",
		Base:        *config_lib.DefaultBaseConfig(),
		GenCodeHost: "localhost:9093",
		JWTSecret:   os.Getenv("jwt_secret"),
	}
}
