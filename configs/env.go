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
	USERGRPCHost string
	USERGRPCPort string
	JWTSecretKey string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		DBUser:       os.Getenv("DBUSER"),
		DBPassword:   os.Getenv("DBPASSWORD"),
		DBName:       os.Getenv("DBNAME"),
		DBHost:       os.Getenv("DBHOST"),
		DBPort:       os.Getenv("DBPORT"),
		USERGRPCHost: os.Getenv("USERGRPCHOST"),
		USERGRPCPort: os.Getenv("USERGRPCPORT"),
		JWTSecretKey: os.Getenv("JWTSECRET"),
	}
}
