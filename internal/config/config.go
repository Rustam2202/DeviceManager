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
	Logger   *logger.LoggerConfig
	Database *database.MongoDbConfig
	Server   *server.ServerHTTPConfig
	Kafka    *kafka.KafkaConfig
}

func MustLoadConfig() *Config {
	var cfg Config
	var err error
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.AddConfigPath(*path)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}

	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg.Server)
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg.Database)
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg.Kafka)
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
