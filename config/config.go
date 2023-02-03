package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetAppName() string
}

type config struct {
	viper *viper.Viper
}

func NewConfig() (Config, error) {
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigName(".env")

	v.AddConfigPath(".")
	v.AddConfigPath("..")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &config{
		viper: v,
	}, nil
}

func (c *config) GetAppName() string {
	return c.viper.GetString("APP_NAME")
}
