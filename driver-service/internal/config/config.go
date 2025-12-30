package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MongoConfig struct {
	ConnString string `mapstructure:"connection_string"`
	Database   string `mapstructure:"database"`
}

type ServiceConfig struct {
	Mongo MongoConfig `mapstructure:"mongo"`
	Port  int         `mapstructure:"port"`
}

func Load() (ServiceConfig, error) {
	var config ServiceConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config/")

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil
}
