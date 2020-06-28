package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type config struct {
	Debug        bool   `mapstructure:"debug"`
	Port         int    `mapstructure:"port"`
	DatabaseType string `mapstructure:"database_type"`
	Database     string `mapstructure:"database"`
}

// default config
var Config *config = &config{
	Debug: true,
	Port:  8000,
}

func Parse() (*config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("VIPER")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(Config)
	if err != nil {
		return nil, err
	}
	log.Debug().
		Interface("config", Config).
		Msg("load config success")

	return Config, nil
}

func Get() *config {
	return Config
}
