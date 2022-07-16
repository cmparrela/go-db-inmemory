package main

import (
	"context"
	"github.com/cmparrela/go-db-inmemory/cache"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	cacher := cache.NewCacher()

	for i := 1; i < 1000000; i++ {
		expiration := time.Duration(rand.Intn(10)) * time.Second

		key := uuid.New().String()
		cacher.Set(key, "teste", expiration)
	}

	<-ctx.Done()
}
