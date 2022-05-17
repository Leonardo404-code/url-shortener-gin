package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	StoreService = &StorageService{}
	Ctx          = context.Background()
)

func InitializeStore() *StorageService {
	dsn := viper.GetString("DATABASE_URL")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatalf("Error init redis: %v", err)
	}

	StoreService.redisClient = redisClient

	return StoreService
}
