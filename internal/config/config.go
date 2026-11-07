package config

import (
	"os"

	"github.com/joho/godotenv"
)

// config holds the application-wide configuration values loaded from the environment.
type Config struct {
	Port string
	Env  string
}

// mustload loads environment variables and validates that required parameters
// are present. it uses a 'must' prefix convention to indicate that a failure
// here is fatal and should panic immediately on startup.
func MustLoad() Config {
	// load variables from a local .env file into the environment.
	// in production environments, this usually fails silently if no .env file is found,
	// relying instead on environment variables injected directly by the host.
	godotenv.Load()

	// retrieve and validate the server port.
	// if missing, the application cannot start correctly, so it panics.
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	// retrieve and validate the runtime environment (e.g., development, production).
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	// return the validated configuration struct.
	return Config{
		Port: port,
		Env:  env,
	}
}
