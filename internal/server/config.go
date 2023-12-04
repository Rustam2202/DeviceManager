package server

type ServerHTTPConfig struct {
	Host string //`mapstructure:"SERVER_HOST"`
	Port int    //`mapstructure:"SERVER_PORT"`
}
