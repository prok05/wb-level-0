package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PublicHost string
	ServerPort string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string

	KafkaURL   string
	KafkaTopic string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "postgres"),
		ServerPort: getEnv("SERVER_PORT", ":8000"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBAddress: fmt.Sprintf("%s:%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5432")),
		DBName:     getEnv("DB_NAME", "l0-db"),
		KafkaURL:   getEnv("KAFKA_URL", "localhost:9092"),
		KafkaTopic: getEnv("KAFKA_TOPIC", "orders"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
