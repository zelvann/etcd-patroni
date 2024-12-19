package config

import "github.com/spf13/viper"

type appEnv struct {
	ApiPort        string `mapstructure:"API_PORT"`
	MinioEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `mapstructure:"MINIO_SECRET_KEY"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
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
