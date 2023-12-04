package config

import (
	"device-manager/internal/database"
	"device-manager/internal/kafka"
	"device-manager/internal/logger"
	"device-manager/internal/server"
	"flag"
	// "strings"

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

	// read from config.yml
	viper.Reset()
	viper.AddConfigPath(*path)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}

	// read from app.env
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("SERVER_HOST", &cfg.Server.Host); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("SERVER_PORT", &cfg.Server.Port); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("DATABASE_HOST", &cfg.Database.Host); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("DATABASE_PORT", &cfg.Database.Port); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("DATABASE_NAME", &cfg.Database.Name); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("KAFKA_BROKERS", &cfg.Kafka.Brokers); err != nil {
		panic(err.Error())
	}
	if err = viper.UnmarshalKey("KAFKA_GROUP", &cfg.Kafka.Group); err != nil {
		panic(err.Error())
	}

	return &cfg
}
