package session

import (
	"context"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

var (
	redisPass = strings.TrimSpace(os.Getenv("REDIS_PASS"))
	redisPort = strings.TrimSpace(os.Getenv("REDIS_PORT"))
	redisHost = strings.TrimSpace(os.Getenv("REDIS_HOST"))
)

type Storage struct {
	rdb *redis.Client
}

// Creates a new session storage
func (s *Storage) Create() {
	options := &redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	}

	s.rdb = redis.NewClient(options)
}

// Gets a value from the session storage
func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	return s.rdb.Get(ctx, key).Result()
}

// Sets a value in the session storage
func (s *Storage) Set(ctx context.Context, key string, value string) error {
	return s.rdb.Set(ctx, key, value, 0).Err()
}

// Deletes a value from the session storage
func (s *Storage) Delete(ctx context.Context, key string) error {
	return s.rdb.Del(ctx, key).Err()
}
