package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const VisitorCountKey = "visitors"

type VisitorStore struct {
	c *redis.Client
}

func NewVisitorStore(redisHost string) *VisitorStore {
	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
	return &VisitorStore{client}
}

func (vs *VisitorStore) Increment() int {
	vs.c.Incr(context.Background(), VisitorCountKey)
	return vs.Visitors()
}

func (vs *VisitorStore) Visitors() int {
	i, err := vs.c.Get(context.Background(), VisitorCountKey).Int()
	if err != nil {
		return 0
	}
	return i
}
