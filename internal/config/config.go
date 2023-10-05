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
	viper.AddConfigPath(*path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}
	viper.AutomaticEnv()

	viper.Reset()
	viper.AddConfigPath(*path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("server")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}
	cfg.ServerHTTPConfig.Host = viper.GetString("SERVER_HTTP_CONFIG_HOST")
	cfg.ServerHTTPConfig.Port = viper.GetInt("SERVER_HTTP_CONFIG_PORT")
	cfg.DatabaseConfig.Host = viper.GetString("DATABASE_CONFIG_HOST")
	cfg.DatabaseConfig.Port = viper.GetInt("DATABASE_CONFIG_PORT")
	cfg.DatabaseConfig.Name = viper.GetString("DATABASE_CONFIG_NAME")
	cfg.KafkaConfig.Brokers = viper.GetStringSlice("KAFKA_CONFIG_BROKERS")
	cfg.KafkaConfig.Group = viper.GetString("KAFKA_CONFIG_GROUP")

	return &cfg
}
