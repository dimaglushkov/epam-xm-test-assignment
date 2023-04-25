package internal

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName    string `env:"APP_NAME,required"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	AppVersion string `env:"APP_VERSION,required"`
	AppMode    string `env:"APP_MODE" envDefault:"debug"`

	DSN string `env:"DB_DSN,required"`
}

// NewConfig returns app configuration of type Config from the given config file path.
func NewConfig() (*Config, error) {
	//err := godotenv.Load()
	//if err != nil {
	//	return nil, err
	//}
	cfg := Config{}
	err := env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
