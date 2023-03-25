package config

import "github.com/thesisK19/buildify/library/conf"

type Config struct {
	conf.Base `mapstructure:",squash"`

	MongoDB   string `mapstructure:"mongo_db"`
	ServiceDB string `mapstructure:"service_db"`
}

func loadDefaultConfig() *Config {
	return &Config{
		MongoDB:   "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net/user_service",
		ServiceDB: "user_service",
		Base:      *conf.DefaultBaseConfig(),
	}
}
