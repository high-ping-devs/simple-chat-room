package auth

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisPass = strings.TrimSpace(os.Getenv("REDIS_PASS"))
	redisPort = strings.TrimSpace(os.Getenv("REDIS_PORT"))
	redisHost = strings.TrimSpace(os.Getenv("REDIS_HOST"))
)

type WhiteList struct {
	rdb *redis.Client
}

// Creates a new whitelist
func (wl *WhiteList) Create() {
	options := &redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	}

	wl.rdb = redis.NewClient(options)
}

// Closes the whitelist
func (wl *WhiteList) Close() {
	wl.rdb.Close()
}

func (wl *WhiteList) Expire(ctx context.Context, key string, minutesToExpire int) error {
	return wl.rdb.Expire(ctx, key, time.Duration(minutesToExpire)*time.Minute).Err()
}

// Gets a value from the whitelist
func (wl *WhiteList) Get(ctx context.Context, key string) (string, error) {
	return wl.rdb.Get(ctx, key).Result()
}

// Gets a hash value from the whitelist
func (wl *WhiteList) HGet(ctx context.Context, key string, field string) (string, error) {
	return wl.rdb.HGet(ctx, key, field).Result()
}

// Sets a hash value in the whitelist
func (wl *WhiteList) HSet(ctx context.Context, key string, field string, value string) error {
	return wl.rdb.HSet(ctx, key, field, value).Err()
}

// Gets all hash values from the whitelist
func (wl *WhiteList) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return wl.rdb.HGetAll(ctx, key).Result()
}

// Deletes a hash value from the whitelist
func (wl *WhiteList) HDel(ctx context.Context, key string, fields ...string) error {
	return wl.rdb.HDel(ctx, key, fields...).Err()
}

// Sets a value in the whitelist
func (wl *WhiteList) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return wl.rdb.Set(ctx, key, value, ttl).Err()
}

// Deletes a value from the whitelist
func (wl *WhiteList) Delete(ctx context.Context, key string) error {
	return wl.rdb.Del(ctx, key).Err()
}
