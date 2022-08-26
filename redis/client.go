package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const pingTimeout = 5 * time.Second

var ErrKeyNotFound = errors.New("key not found")

type Redis struct {
	client *redis.Client
}

func NewClient(ctx context.Context, url string) (*Redis, error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	client := redis.NewClient(options)

	pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()

	if err = client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &Redis{client: client}, nil
}

// Set sets a vlue by the key.
func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	if _, err = r.client.Set(ctx, key, data, expiration).Result(); err != nil {
		return fmt.Errorf("failed to set a value by key - %s: %w", key, err)
	}

	return nil
}

// Get reads a value by the key.
func (r *Redis) Get(ctx context.Context, key string, receiver interface{}) error {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrKeyNotFound
		}

		return fmt.Errorf("failed to read a value by key - %s: %w", key, err)
	}

	if err = json.Unmarshal([]byte(value), receiver); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

// Delete deletes a value by the key.
func (r *Redis) Delete(ctx context.Context, key string) error {
	if _, err := r.client.Del(ctx, key).Result(); err != nil {
		return fmt.Errorf("failed to delete a value by key - %s: %w", key, err)
	}

	return nil
}

func (r *Redis) IsAvailable(ctx context.Context) bool {
	return r.client.Ping(ctx).Err() == nil
}

func (r *Redis) Reconnect(ctx context.Context, host string) error {
	options, err := redis.ParseURL(host)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	client := redis.NewClient(options)
	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to ping client: %w", err)
	}

	r.client = client

	return nil
}
