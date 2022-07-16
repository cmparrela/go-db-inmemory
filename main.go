package main

import (
	"context"
	"github.com/cmparrela/go-db-inmemory/cache"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	cacher := cache.NewCacher()

	// Create items in cache
	for i := 1; i < 100000; i++ {
		current := time.Now()
		expiration := current.Add(time.Duration(rand.Intn(10)) * time.Second)

		key := uuid.New().String()
		cacher.Set(key, "teste", expiration)
	}

	var m sync.Mutex
	go cacher.Garbageman(&m)

	<-ctx.Done()
}
