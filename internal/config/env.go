package config


import (

	"log"

	"os"


	"github.com/joho/godotenv"

)


func LoadEnv() {

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error loading .env file")

	}


	log.Println("Environment variables loaded")

	log.Println("DB Host:", os.Getenv("DB_HOST"))

}