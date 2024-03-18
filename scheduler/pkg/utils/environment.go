package utils

import (
	"github.com/caarlos0/env/v6"
)

var EnvironmentVars Environment

type Environment struct {
	DatabaseHost            string `env:"DB_HOST,notEmpty"`
	DatabasePort            string `env:"DB_PORT,notEmpty"`
	DatabaseUser            string `env:"DB_USER,notEmpty"`
	DatabasePassword        string `env:"DB_PASSWORD,notEmpty"`
	DatabaseName            string `env:"DB_NAME,notEmpty"`
	DatabaseSslMode         string `env:"DB_SSL_MODE,notEmpty"`
	DatabaseMaxConn         int    `env:"DB_MAX_CONNECTIONS,notEmpty"`
	DatabaseMaxIdleConn     int    `env:"DB_MAX_IDLE_CONNECTIONS,notEmpty"`
	DatabaseMaxLifetimeConn int    `env:"DB_MAX_LIFETIME_CONNECTIONS,notEmpty"`
	RedisHost               string `env:"REDIS_HOST,notEmpty"`
	RedisPort               string `env:"REDIS_PORT,notEmpty"`
	RedisPassword           string `env:"REDIS_PASSWORD,notEmpty"`
	RedisDbNumber           int    `env:"REDIS_DB_NUMBER,notEmpty"`
	WebhookApiUrl           string `env:"WEBHOOK_API_URL,notEmpty"`
	WebhookApiKey           string `env:"WEBHOOK_API_KEY,notEmpty"`
	CronJobSchedule         string `env:"CRON_JOB_SCHEDULE,notEmpty"`
}

func ParseEnvironmentVariables() error {
	// load env variables to struct
	if err := env.Parse(&EnvironmentVars); err != nil {
		return err
	}
	return nil
}
