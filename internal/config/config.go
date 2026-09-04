package config

import (
	"github.com/joho/godotenv"
	"os"
)

// C :- capital means export for public
type Config struct {
	Port string
	Env  string
}

// must pattern :- if the function fails it will panic and stop the program.
func MustLoad() Config {
	// Load the .env file (optional)
	godotenv.Load()

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // default env
		// panic("APP_ENV required but not found in .env") :- this is used for panic
	}

	return Config{
		Port: port,
		Env:  env,
	}
}
