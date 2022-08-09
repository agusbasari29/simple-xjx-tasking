package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConnection struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

func DbConfig() DbConnection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file!")
	}

	dbConfig := DbConnection{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}

	return dbConfig
}
