package redis

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/Kpatoc452/Dceorator/models"
	"github.com/redis/go-redis/v9"
)

type Redis interface{
	Get(id int) (models.Info, error)
}


type redisCache struct {
	rds *redis.Client
}

func New() *redisCache{
	return &redisCache{
		rds: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
	}
}

func(r *redisCache) Get(id int) (models.Info, error) {
	var info models.Info
	val, err := r.rds.Get(context.Background(),  strconv.Itoa(id)).Result()
    if err != nil {
		log.Println("Redis Get. Error")
        return info, err
    } else {
		err = json.Unmarshal([]byte(val), &info)
		if err != nil{ 
			log.Println("json unmrarshall. Error")
		}
		return info, err
	}
}

func(r *redisCache) set(id int, info models.Info) error { 
	info_b, err := json.Marshal(info)
	if err != nil{
		log.Println("Redis. Cannot to SET data.")
		return err
	}
	err = r.rds.Set(context.Background(), strconv.Itoa(id), info_b, 0).Err()
	return err
}

