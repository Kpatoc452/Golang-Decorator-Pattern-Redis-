package service

import (
	"github.com/Kpatoc452/Dceorator/models"
	"github.com/Kpatoc452/Dceorator/storage/postgres"
	"github.com/Kpatoc452/Dceorator/storage/redis"
)

type info struct{
	psql postgres.DB
	rds redis.Redis
}

func newInfo(psql postgres.DB, rds redis.Redis) *info{
	return &info{
		psql: psql,
		rds: rds,
	}
}

func(s *info) GetInfo(id int) (models.Info, error){
	return  
}

func(s *info) AddInfo(info string) error {

}

func(s *info) GetAllInfo() ([]string, error){

}

func(s *info) DeleteInfo(id int) error { 

}


