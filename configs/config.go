package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func NewConfig(pathToConfig string) (*Config, error) {
	newConf := &Config{}

	viper.AddConfigPath(pathToConfig)

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(newConf); err != nil {
		return nil, err
	}
	return newConf, nil
}
