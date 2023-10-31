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
	LoggerConfig *logger.LoggerConfig
	Database     *database.MongoDbConfig
	Server       *server.ServerHTTPConfig
	KafkaConfig  *kafka.KafkaConfig
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

	// fmt.Println("Server.Host:", cfg.Server.Host)
	// fmt.Println("Server.Port:", cfg.Server.Port)
	// fmt.Println("Database.Host:", cfg.Database.Host)
	// fmt.Println("Database.Port:", cfg.Database.Port)
	// fmt.Println("SERVER_HOST:", viper.Get("SERVER_HOST"))
	// fmt.Println("SERVER_PORT:", viper.GetInt("SERVER_PORT"))
	// fmt.Println("DATABSE_HOST:", viper.GetString("DATABASE_HOST"))
	// fmt.Println("DATABASE_PORT", viper.GetInt("DATABASE_PORT"))
	// fmt.Println("DATABASE_NAME", viper.GetString("DATABASE_NAME"))
	// fmt.Println("KAFKA_BROKERS", viper.GetStringSlice("KAFKA_BROKERS"))
	// fmt.Println("KAFKA_GROUP", viper.GetString("KAFKA_GROUP"))

	return &cfg
}
