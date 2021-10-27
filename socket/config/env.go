package config

import "github.com/ilyakaznacheev/cleanenv"

type Env struct {
	REDIS_URI         string `env:"REDIS_URI" env-default:"http://localhost:6379"`
	REDIS_PASSWORD    string `env:"REDIS_PASSWORD" env-default:""`
	REDIS_DB          int    `env:"REDIS_DB" env-default:"0"`
	AUTH_SERVICE_HOST string `env:"AUTH_SERVICE_HOST" env-default:"http://localhost:4000"`
	WS_AUTH_PATH      string `env:"WS_AUTH_PATH" env-default:"/authentication/ws"`
	APP_PORT          string `env:"APP_PORT" env-default:":8080"`
}

func GetEnv() (*Env, error) {
	var env Env
	if err := cleanenv.ReadEnv(&env); err != nil {
		return nil, err
	}
	return &env, nil
}
