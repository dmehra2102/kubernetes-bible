package config

import "os"

type Config struct {
	Broker string
	Topic  string
	Group  string
}

func Load() *Config {
	return &Config{
		Broker: os.Getenv("KAFKA_BROKER"),
		Topic : os.Getenv("KAFKA_TOPIC"),
		Group: os.Getenv("KAFKA_GROUP"),
	}
}