package config

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Load()(*Config, error) {
	c := loadDefaultConfig()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config file, load default.")

		configBuffer, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}

		err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
		if err != nil {
			return nil, err
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
