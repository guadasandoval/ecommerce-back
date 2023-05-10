package keyvalue

import (

	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

// Initialize function inicializa la db en memoria
func Initialize(host string, port int, database int, password string) {
	if redisClient != nil {
		panic("KeyValue already initialized")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		DB:       database,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Redis connection error: " + err.Error())
	}

	redisClient = client
}
