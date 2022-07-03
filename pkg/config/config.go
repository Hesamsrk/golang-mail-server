package config

import "github.com/spf13/viper"

type Config struct {
	DB_PASSWORD    string `mapstructure:"DB_PASSWORD"`
	JWT_SECRET     string `mapstructure:"JWT_SECRET"`
	EMAIL_ADDRESS  string `mapstructure:"EMAIL_ADDRESS"`
	EMAIL_PASSWORD string `mapstructure:"EMAIL_PASSWORD"`
}

var LocalConfig = &Config{
	DB_PASSWORD:    "",
	JWT_SECRET:     "",
	EMAIL_ADDRESS:  "",
	EMAIL_PASSWORD: "",
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
