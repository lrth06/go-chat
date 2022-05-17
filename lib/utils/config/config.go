package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lrth06/go-chat/lib/structs"
)

func GetConfig() (structs.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := structs.Config{}
	config.Port = os.Getenv("PORT")
	config.AppEnv = os.Getenv("APP_ENV")
	config.MongoURI = os.Getenv("MONGO_URI")
	config.DBName = os.Getenv("DB_NAME")
	config.TokenSecret = os.Getenv("JWT_SECRET")
	return config, err
}
