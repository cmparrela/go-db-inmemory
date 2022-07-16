package cache

import (
	"fmt"
)

type Repository interface {
	List() map[string]*Cache
	Get(key string) (*Cache, error)
	Create(cache *Cache)
	Delete(cache *Cache)
}

type repository struct {
	data map[string]*Cache
}

func NewRepository() Repository {
	var mp = map[string]*Cache{}
	return &repository{data: mp}
}

func (r *repository) List() map[string]*Cache {
	return r.data
}

func (r *repository) Get(key string) (*Cache, error) {
	if r.data[key] == nil {
		return nil, fmt.Errorf("key not found")
	}
	return r.data[key], nil
}

func (r *repository) Create(cache *Cache) {
	r.data[cache.Key] = cache
}

func (r *repository) Delete(cache *Cache) {
	delete(r.data, cache.Key)
}
