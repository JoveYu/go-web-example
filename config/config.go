package config

import (
	"github.com/spf13/viper"
)

type config struct {
	Debug bool `yaml:"debug" json:"debug" toml:"debug"`
	Port  int  `yaml:"port" json:"port" toml:"port"`
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
	return Config, nil
}

func Get() *config {
	return Config
}
