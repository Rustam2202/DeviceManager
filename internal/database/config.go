package database

type MongoDbConfig struct {
	Host string //`mapstructure:"DATABASE_HOST"`
	Port int    //`mapstructure:"DATABASE_PORT"`
	Name string //`mapstructure:"DATABASE_NAME"`
}
