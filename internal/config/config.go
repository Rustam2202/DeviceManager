package config

import (
	"device-manager/internal/database"
	"device-manager/internal/kafka"
	"device-manager/internal/logger"
	"device-manager/internal/server"
	"flag"

	"github.com/spf13/viper"
)

type Config struct {
	LoggerConfig     *logger.LoggerConfig
	DatabaseConfig   *database.MongoDbConfig
	ServerHTTPConfig *server.ServerHTTPConfig
	KafkaConfig      *kafka.KafkaConfig
}

func MustLoadConfig() *Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}
	return &cfg
}
