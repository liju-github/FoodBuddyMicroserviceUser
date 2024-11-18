package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser       string
	DBPassword   string
	DBName       string
	DBHost       string
	DBPort       string
	GRPCPort     string
	JWTSecretKey string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		GRPCPort:     os.Getenv("GRPC_PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET"),
	}
}
