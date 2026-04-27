package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DB         string
	Secret_jwt string

	Db_host string
	Db_user string
	Db_name string
	Db_pass string
	Db_port string
}

func MustLoad() (*Config, error) {
	err := godotenv.Load()
	if err!=nil {
		return nil, fmt.Errorf("failed to load .env file")
	}

	log.Println("config loaded")

	return &Config{
		Port: os.Getenv("PORT"),
		DB: os.Getenv("DATABASE_URL"),
		Secret_jwt: os.Getenv("SECRET_JWT"),
		Db_host: os.Getenv("DB_HOST"),
		Db_user: os.Getenv("DB_USER"),
		Db_pass: os.Getenv("DB_PASSWORD"),
		Db_name: os.Getenv("DB_NAME"),
		Db_port: os.Getenv("DB_PORT"),
	},nil
}