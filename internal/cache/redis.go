package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	return &RedisCache{client: rdb}
}
func (c *RedisCache) Get(key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("[CACHE] MISS", key)
		return "", err
	}
	log.Println("[CACHE] HIT", key)
	return val, nil
}
func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	log.Println("[CACHE] SET", key)
	log.Println("[CACHE] DATA =", value)
	return c.client.Set(ctx, key, value, ttl).Err()
}
func (c *RedisCache) Del(key string) error {
	log.Println("[CACHE] DEL", key)
	return c.client.Del(ctx, key).Err()
}
