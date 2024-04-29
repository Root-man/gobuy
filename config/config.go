package config

import "github.com/spf13/viper"

type BotConfig struct {
	token string
}

func (bc *BotConfig) GetToken() string {
	return bc.token
}

func Load() (*BotConfig, error) {
	viper.SetEnvPrefix("GOBUY")
	viper.AutomaticEnv()

	return &BotConfig{token: viper.GetString("TOKEN")}, nil
}
