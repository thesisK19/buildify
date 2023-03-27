package conf

import (
	"time"

	"github.com/thesisK19/buildify/library/log"
	server "github.com/thesisK19/buildify/library/server"
)

// deploy env.
const (
	DeployEnvDev  = "dev"
	DeployEnvProd = "prod"
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
			Port: 10443,
		},
		HTTP: server.Listen{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}
}

// Config ...
type Base struct {
	Env string     `json:"env" mapstructure:"env"`
	Log log.Config `json:"log" mapstructure:"log"`
	// LogLevel int `json:"log_level" mapstructure: "log_level"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

func (b Base) IsDevelopment() bool {
	return b.Env == DeployEnvDev
}

func DefaultBaseConfig() *Base {
	return &Base{
		Env: DeployEnvDev,
		// LogLevel: 2,
		Log:    log.DefaultConfig(),
		Server: DefaultServerConfig(),
	}
}
