package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	MongoDBUri   string `mapstructure:"MongoDB_URI"`
	Host         string `mapstructure:"HOST"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
}

func LoadEnv(path string) (Env, error) {
	var env Env
	if path == "" {
		path = "."
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&env)
	return env, err
}
