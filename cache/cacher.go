package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cacher interface {
	List() map[string]*Cache
	Set(key string, value string, expiration time.Duration)
	Get(key string) Cache
	Delete(key string)
}

type cacher struct {
	repository Repository
}

func NewCacher() Cacher {
	return &cacher{
		repository: NewRepository(&sync.Mutex{}),
	}
}

func (c cacher) List() map[string]*Cache {
	return c.repository.List()
}

func (c cacher) Set(key string, value string, expiration time.Duration) {
	cache := Cache{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	c.repository.Create(&cache)

	// https://gobyexample.com/timeouts
	go func() {
		select {
		case <-time.After(expiration):
			c.Delete(key)
			fmt.Println("cache deleted ", key)
			return
		}
	}()
}

func (c cacher) Get(key string) Cache {
	return c.repository.Get(key)
}

func (c cacher) Delete(key string) {
	c.repository.Delete(key)
}
