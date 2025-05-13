package config

import (
	"os"
	"sync"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	GRPCPort         string
	HTTPPort         string
	NATSURL          string
	UserServiceHost  string
	UserServicePort  string
	MangaServiceHost string
	MangaServicePort string
	OrderServiceHost string
	OrderServicePort string
}

var (
	once     sync.Once
	instance *Config
)

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			DBHost:           getEnv("DB_HOST", "localhost"),
			DBPort:           getEnv("DB_PORT", "5432"),
			DBUser:           getEnv("DB_USER", "postgres"),
			DBPassword:       getEnv("DB_PASSWORD", "postgres"),
			DBName:           getEnv("DB_NAME", "mangaverse"),
			GRPCPort:         getEnv("GRPC_PORT", "50051"),
			HTTPPort:         getEnv("HTTP_PORT", "8080"),
			NATSURL:          getEnv("NATS_URL", "nats://localhost:4222"),
			UserServiceHost:  getEnv("USER_SERVICE_HOST", "localhost"),
			UserServicePort:  getEnv("USER_SERVICE_PORT", "50051"),
			MangaServiceHost: getEnv("MANGA_SERVICE_HOST", "localhost"),
			MangaServicePort: getEnv("MANGA_SERVICE_PORT", "50052"),
			OrderServiceHost: getEnv("ORDER_SERVICE_HOST", "localhost"),
			OrderServicePort: getEnv("ORDER_SERVICE_PORT", "50053"),
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
