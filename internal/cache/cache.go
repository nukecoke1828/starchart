package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

type Cache struct {
	codec *cache.Codec
}

func New(redisURL string) (*Cache, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)

	return &Cache{
		codec: &cache.Codec{
			Redis: client,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
	}, nil
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.codec.Set(&cache.Item{
		Ctx:    ctx,
		Key:    key,
		Value:  value,
		TTL:    ttl,
	})
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	return c.codec.Get(ctx, key, value)
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.codec.Delete(ctx, key)
}