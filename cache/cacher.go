package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cacher interface {
	List() map[string]*Cache
	Set(key string, value string, expiration time.Time)
	Get(key string) (*Cache, error)
	Delete(key string)
	Garbageman(m *sync.Mutex)
}

type cacher struct {
	repository Repository
	done       chan bool
}

func NewCacher() Cacher {
	repository := NewRepository()
	return &cacher{repository: repository}
}

func (c cacher) Garbageman(m *sync.Mutex) {
	for {
		garbages := c.List()

		m.Lock()
		for i := range garbages {
			expiration := garbages[i].Expiration

			go func(expiration time.Time, i string, m *sync.Mutex) {
				m.Lock()
				now := time.Now()
				if now.After(expiration) {
					c.Delete(i)
					fmt.Println("cache deleted ", i)
				}
				m.Unlock()

			}(expiration, i, m)
		}
		m.Unlock()
	}

}

func (c cacher) List() map[string]*Cache {
	return c.repository.List()
}

func (c cacher) Set(key string, value string, expiration time.Time) {
	cache := Cache{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	c.repository.Create(&cache)
}

func (c cacher) Get(key string) (*Cache, error) {
	return c.repository.Get(key)
}

func (c cacher) Delete(key string) {
	cache, _ := c.repository.Get(key)
	c.repository.Delete(cache)
}
