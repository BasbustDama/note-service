package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseDsn string `mapstructure:"DATABASE_DSN"`
	HttpPort    string `mapstructure:"HTTP_PORT"`
}

func MustGetConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}

	return config
}
