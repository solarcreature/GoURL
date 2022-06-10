package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type StorageService struct {
	redisClient *redis.redisClient
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"",
		DB:			0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v",err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}",pong)
	storeService.redisClient = redisClient

	return storeService
}

func StoreMap(shortURL string, originalURL string, userID string) {
	err := storeService.redisClient.Set(ctx,shortURL,originalURL,CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed storing key url | Error: %v - shortURL: %s - originalURL: %s\n",err,shortURL,originalURL))

	}

}

func RetrieveOriginalURL(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx,shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve the initial url | Error: %v - shortURL: %s\n",err,shortURL))
	}

	return result
}
