package config

import "github.com/ilyakaznacheev/cleanenv"

type Env struct {
	REDIS_URI              string `env:"REDIS_URI" env-default:"localhost:6379"`
	REDIS_PASSWORD         string `env:"REDIS_PASSWORD" env-default:""`
	REDIS_DB               int    `env:"REDIS_DB" env-default:"0"`
	AUTH_SERVICE_HOST      string `env:"AUTH_SERVICE_HOST" env-default:"http://localhost:4000"`
	WS_AUTH_PATH           string `env:"WS_AUTH_PATH" env-default:"/authentication/ws"`
	APP_PORT               string `env:"APP_PORT" env-default:":8080"`
	INIT_WOKER_NUM         int    `env:"INIT_WOKER_NUM" env-default:"1"`
	MAX_WORKER_NUM         int    `env:"MAX_WORKER_NUM" env-default:"3"`
	TASK_QUEUE_SIZE        int    `env:"TASK_QUEUE_SIZE" env-default:"5"`
	KAFKA_NETWORK_PROTOCOL string `env:"KAFKA_NETWORK_PROTOCOL" env-default:"tcp"`
	KAFKA_URI              string `env:"KAFKA_URI" env-default:"localhost:9092"`
	KAFKA_TOPIC            string `env:"KAFKA_TOPIC" env-default:"message"`
}

func GetEnv() (*Env, error) {
	var env Env
	if err := cleanenv.ReadEnv(&env); err != nil {
		return nil, err
	}
	return &env, nil
}
