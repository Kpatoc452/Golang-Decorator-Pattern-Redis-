package main

import (
	"context"
	"log"
	"net/http"

	// "github.com/jackc/pgx/v5"
	// "github.com/joho/godotenv"

	"github.com/Kpatoc452/Dceorator/config"
	"github.com/Kpatoc452/Dceorator/storage/postgres"
	"github.com/Kpatoc452/Dceorator/storage/redis"
)

var ctx = context.Background()



func main(){
	cfg, err := config.New()
	if err != nil{
		log.Fatal(err)
	}

	db, err := postgres.New(cfg)
	if err != nil{
		log.Fatal(err)
	}

	redis := redis.New()
	

	http.HandleFunc("/info", )

}



