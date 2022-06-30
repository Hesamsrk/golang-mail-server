package config

import "github.com/spf13/viper"

type Config struct {
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
}

var LocalConfig = &Config{
	DB_PASSWORD: "password",
	JWT_SECRET:  "secret",
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
