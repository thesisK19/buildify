package config

import (
	"github.com/thesisK19/buildify/library/log"
	server_lib "github.com/thesisK19/buildify/library/server"
)

type Config struct {
	HTTP      server_lib.Listen `json:"http" mapstructure:"http"`
	MongoDB   string            `json:"mongo_db" mapstructure:"mongo_db"`
	Log       log.Config        `json:"log" mapstructure:"log"`
	ServiceDB string            `json:"service_db" mapstructure:"service_db"`
}

func loadDefaultConfig() *Config {
	return &Config{
		MongoDB:   "mongodb+srv://thesis:thesisK19@thesis.kzystcv.mongodb.net",
		ServiceDB: "file_management_service",
		HTTP: server_lib.Listen{
			Host: "0.0.0.0",
			Port: 80,
		},
	}
}
