package config

import (
	"github.com/thesisK19/buildify/library/log"
	"github.com/thesisK19/buildify/library/server"
)

type Config struct {
	HTTP      server.Listen `json:"http" mapstructure:"http" yaml:"http"`
	MongoDB   string        `mapstructure:"mongo_db"`
	Log       log.Config    `json:"log" mapstructure:"log"`
	ServiceDB string        `mapstructure:"service_db"`
}

func loadDefaultConfig() *Config {
	return &Config{
		MongoDB:   "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net/file_management_service",
		ServiceDB: "file_management_service",
		HTTP: server.Listen{
			Host: "0.0.0.0",
			Port: 80,
		},
	}
}