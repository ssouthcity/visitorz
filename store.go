package main

import (
	"os"

	"github.com/ssouthcity/visitorz/inmem"
	"github.com/ssouthcity/visitorz/redis"
)

type VisitorStore interface {
	Increment() int
	Visitors() int
}

func visitorStoreFactory() VisitorStore {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost != "" {
		return redis.NewVisitorStore(redisHost)
	}

	return inmem.NewVisitorStore()
}
