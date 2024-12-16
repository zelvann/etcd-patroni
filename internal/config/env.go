package config

import "github.com/spf13/viper"

type appEnv struct {
	ApiPort string `mapstructure:"PORT"`
}

func LoadEnv() *appEnv {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	var res appEnv
	if err := viper.Unmarshal(&res); err != nil {
		panic(err)
	}

	return &res
}
