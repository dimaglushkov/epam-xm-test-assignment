package internal

import (
	"github.com/caarlos0/env/v6"
)

// Config represents applications config loaded from environment variables with default values
type Config struct {
	AppName    string `env:"APP_NAME,required"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	AppVersion string `env:"APP_VERSION,required"`
	AppMode    string `env:"APP_MODE" envDefault:"debug"`
	AppSignKey string `env:"APP_SIGN_KEY,required"`

	DSN                  string `env:"DB_DSN,required"`
	DBApplyMigrations    int    `env:"DB_APPLY_MIGRATIONS" envDefault:"1"`
	DBMaxPoolSize        int    `env:"DB_MAX_POOL_SIZE" envDefault:"1"`
	DBConnAttempts       int    `env:"DB_CONN_ATTEMPTS" envDefault:"10"`
	DBConnTimeoutSeconds int    `env:"DB_CONN_TIMEOUT_SECONDS" envDefault:"1"`

	KafkaBrokers []string `env:"KAFKA_BROKERS,required"`
	KafkaTopic   string   `env:"KAFKA_TOPIC,required"`
}

// NewConfig returns app configuration of type Config
func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
