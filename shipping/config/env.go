package config

import "github.com/ilyakaznacheev/cleanenv"

type Env struct {
	KAFKA_NETWORK_PROTOCOL string `env:"KAFKA_NETWORK_PROTOCOL" env-default:"tcp"`
	KAFKA_URI              string `env:"KAFKA_URI" env-default:"localhost:9092"`
	KAFKA_TOPIC            string `env:"KAFKA_TOPIC" env-default:"message"`
	KAFKA_GROUP_CONSUMER   string `env:"KAFKA_GROUP_CONSUMER" env-default:"shippers"`
	KAFKA_PARTITION        int    `env:"KAFKA_PARTITION" env-defalt:0`
	ELASTIC_SEARCH_URI     string `env:"ELASTIC_SEARCH_URI" env-default:"http://localhost:9200"`
}

func GetEnv() (*Env, error) {
	var env Env
	if err := cleanenv.ReadEnv(&env); err != nil {
		return nil, err
	}
	return &env, nil
}
