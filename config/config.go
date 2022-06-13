package config

import (
	"fmt"
	"os"
)

var Config Configuration

func init() {
	LoadEnv()
	Config = Load()
}

type Configuration struct {
	// gRPC Configurations
	Service_HOST string
	Service_PORT string

	// PostgreSQL Configuration
	PostgresHost     string
	PostgresPort     string
	PostgresUsername string
	PostgresPassword string
	PostgresSSLMode  string
	LaptopServiceDB  string
}

func Load() Configuration {
	config := Configuration{}

	config.Service_HOST = anyToString(getEnvOrDefault("SERVICE_HOST", "0.0.0.0"))
	config.Service_PORT = anyToString(getEnvOrDefault("SERVICE_PORT", 5000))

	config.PostgresHost = anyToString(getEnvOrDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = anyToString(getEnvOrDefault("POSTGRES_PORT", 5432))
	config.PostgresUsername = anyToString(getEnvOrDefault("POSTGRES_USERNAME", "postgres"))
	config.PostgresPassword = anyToString((getEnvOrDefault("POSTGRES_PASSWORD", "postgres")))
	config.PostgresSSLMode = anyToString(getEnvOrDefault("POSTGRES_SSL_MODE", "disable"))
	config.LaptopServiceDB = anyToString(getEnvOrDefault("LAPTOP_SERVICE_DB", "laptop_service"))

	return config
}

func anyToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func getEnvOrDefault(key string, default_value interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return default_value
}
