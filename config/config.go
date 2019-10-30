package config

import (
"fmt"
"github.com/spf13/viper"
)

type Config struct {
	Load 		int
	CountryCode []string
	ProxyType	string
}

var Cfg Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Cfg = Config{
		Load:           viper.GetInt("max_load"),
		CountryCode:    viper.GetStringSlice("country_code"),
		ProxyType:      viper.GetString("proxy"),
	}
}
