package cache

import (
	"sync"
)

type Repository interface {
	List() map[string]*Cache
	Get(key string) Cache
	Create(cache *Cache)
	Delete(key string)
}

type repository struct {
	data map[string]*Cache
	mu   *sync.Mutex
}

func NewRepository(mu *sync.Mutex) Repository {
	return &repository{
		data: map[string]*Cache{},
		mu:   mu,
	}
}

func (r *repository) List() map[string]*Cache {
	return r.data
}

func (r *repository) Get(key string) Cache {
	r.mu.Lock()
	result := r.data[key]
	r.mu.Unlock()
	return *result
}

func (r *repository) Create(cache *Cache) {
	r.mu.Lock()
	r.data[cache.Key] = cache
	r.mu.Unlock()
}

func (r *repository) Delete(key string) {
	r.mu.Lock()
	delete(r.data, key)
	r.mu.Unlock()
}
