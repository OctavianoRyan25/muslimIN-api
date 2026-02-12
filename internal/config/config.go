package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type Redis struct {
	Addr     string
	Password string
	DB       uint
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No File ENV Found!")
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func LoadRedis() *Redis {
	return &Redis{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
}
