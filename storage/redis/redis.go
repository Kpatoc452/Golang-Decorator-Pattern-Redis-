package cache

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Kpatoc452/Dceorator/models"
	"github.com/redis/go-redis/v9"
)

type Redis interface{
	Get(id int) (models.Info, error)
}


type redisCache struct {
	rds *redis.Client
}

func NewRedis() *redisCache{
	return &redisCache{
		rds: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
	}
}


func(r *redisCache) Get(id int) (models.Info, error) {
	val, err := r.rds.Get(context.Background(),  string(id)).Result()
    if err != nil {
        
    } else {
		var info models.Info
		err = json.Unmarshal([]byte(val), info)
		return info, err
	}

}



