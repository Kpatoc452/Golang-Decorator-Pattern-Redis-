package middleware

import (
	"github.com/Kpatoc452/Dceorator/models"
	"github.com/Kpatoc452/Dceorator/storage/redis"
)


type GetFromPSQL func(id int) (models.Info, error)

func Caching (f GetFromPSQL, rds redis.Redis) (models.Info, error) {
	var info models.Info
	info, err := rds.Get()
}