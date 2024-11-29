package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetPSQL() string
}

type config struct {
}

func New() (config, error){
	err := godotenv.Load()
	if err != nil {
		log.Println("NewConfig. Error Get .env")
	}
	return config{}, err
}

func(c config) GetPSQL() string{
	return os.Getenv("DATABASE_URL")
}

