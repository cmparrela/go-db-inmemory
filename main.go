package main

import (
	"context"
	"fmt"
	"github.com/cmparrela/go-db-inmemory/cache"
	"github.com/google/uuid"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	cacher := cache.NewCacher()

	// Create items in cache
	for i := 1; i < 1000000; i++ {
		current := time.Now()
		//time.Duration(rand.Intn(10))
		//expiration := current.Add(1 * time.Second)

		key := uuid.New().String()
		fmt.Println("inserindo ", key)
		cacher.Set(key, "teste", current)
	}

	var m sync.Mutex
	go cacher.Garbageman(&m)

	<-ctx.Done()
}
