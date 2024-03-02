package config

import (
	"errors"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type BotConfig struct {
	BotToken string
}

func LoadConfig() (BotConfig, error) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("config/config.yaml")
	if err != nil {
		panic(err)
	}

	switch v := config.Data()["botToken"].(type) {
	case string:
		return BotConfig{BotToken: v}, nil
	default:
		return BotConfig{}, errors.New("invalid bot token in config")
	}

	// fmt.Printf("config data: \n %#v\n", config.Data())
}
