package conf

import (
	"time"

	"github.com/thesisK19/buildify/library/log"
	server "github.com/thesisK19/buildify/library/server"
)

const (
	ContainerStartTimeout = 5 * time.Minute
)

// Config hold http/grpc server config
type ServerConfig struct {
	GRPC server.Listen `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP server.Listen `json:"http" mapstructure:"http" yaml:"http"`
}

// DefaultServerConfig return a default server config
func DefaultServerConfig() ServerConfig {
	//nolint:gomnd
	return ServerConfig{
		GRPC: server.Listen{
			Host: "0.0.0.0",
			Port: 443,
		},
		HTTP: server.Listen{
			Host: "0.0.0.0",
			Port: 80,
		},
	}
}

// Config ...
type Base struct {
	Log    log.Config   `json:"log" mapstructure:"log"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

func DefaultBaseConfig() *Base {
	return &Base{
		Log:    log.DefaultConfig(),
		Server: DefaultServerConfig(),
	}
}
