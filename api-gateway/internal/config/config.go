package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type TokenConfig struct {
	PrivateKeyPath        string `mapstructure:"private_key_path"`
	PublicKeyPath         string `mapstructure:"public_key_path"`
	TokenExpiresInMinutes int    `mapstructure:"token_expires_in_minutes"`
}

type DriverService struct {
	Url string `mapstructure:"url"`
}

type ServiceConfig struct {
	Token         TokenConfig   `mapstructure:"token"`
	Port          int           `mapstructure:"port"`
	DriverService DriverService `mapstructure:"driver-service"`
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
