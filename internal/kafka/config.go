package kafka

type KafkaConfig struct {
	Brokers []string //`mapstructure:"KAFKA_BROKERS"`
	Group   string   //`mapstructure:"KAFKA_GROUP"`
}
