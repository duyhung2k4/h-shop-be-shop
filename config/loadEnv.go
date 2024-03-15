package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	appPort = os.Getenv(APP_PORT)
	dbHost = os.Getenv(DB_HOST)
	dbPort = os.Getenv(DB_PORT)
	dbName = os.Getenv(DB_NAME)
	dbPassword = os.Getenv(DB_PASSWORD)
	dbUser = os.Getenv(DB_USER)
	urlRedis = os.Getenv(URL_REDIS)

	return nil
}
