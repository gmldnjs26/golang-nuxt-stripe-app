package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Cache        *redis.Client
	CacheChannel chan string
)

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string) // acceptor

	go func(ch chan string) {
		for {
			temp := <-ch
			time.Sleep(5 * time.Second)

			Cache.Del(context.Background(), temp)

			fmt.Println("Cache Clear2", temp)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) { // sender
	for _, key := range keys {
		CacheChannel <- key
	}
}
