package internal

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	Environ struct {
		ServerPort int           `required:"true" envconfig:"PORT"`
		JwtSecret  string        `required:"true" envconfig:"SECRET_KEY"`
		JwtExp     time.Duration `required:"true" envconfig:"JWT_EXP"`
		DBHost     string        `required:"true" envconfig:"DB_HOST"`
		DBPort     string        `required:"true" envconfig:"DB_PORT"`
		DBUser     string        `required:"true" envconfig:"DB_USER"`
		DBPass     string        `required:"true" envconfig:"DB_PASS"`
		DBName     string        `required:"true" envconfig:"DB_NAME"`
	}
)

// Get required environment variables
func GetEnv() (*Environ, error) {
	e := new(Environ)

	// Process and validate the required environment variables
	err := envconfig.Process("INST", e)
	if err != nil {
		return nil, err
	}

	return e, nil
}
