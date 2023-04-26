package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	c := loadDefaultConfig()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config file, load default")

		configBuffer, err := json.Marshal(c)
		if err != nil {
			log.Fatal("failed to malshal config", err)
			return nil, err
		}

		err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
		if err != nil {
			log.Fatal("failed to read config by viper", err)
			return nil, err
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatal("failed to unmarshal config by viper", err)
		return nil, err
	}

	fmt.Printf("gen code service: %s\n", c.GenCodeHost)
	return c, nil
}
