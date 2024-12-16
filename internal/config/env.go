package config

import "github.com/spf13/viper"

type appEnv struct {
	ApiPort        string `mapstructure:"API_PORT"`
	MinioEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `mapstructure:"MINIO_SECRET_KEY"`
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
