package cache 

import (
	"errors"
	"sync"
	"time"
)

type ItemCache struct {
	value interface{}
	ttl time.Duration
	createdAt time.Time
}

type Cache struct {
	storage map[string]ItemCache
	mu sync.RWMutex
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]ItemCache),
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	value := c.storage[key]
	if value.createdAt.UnixMicro() < time.Now().UnixMicro() - value.ttl.Microseconds() {
		return ItemCache{}, errors.New("value for this key already exists")
	}
	return value.value, nil
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	existsValue := c.storage[key]
	if existsValue.value != nil {
		return false, errors.New("value for this key already exists")
	}
	c.storage[key] = ItemCache{value: value, ttl: ttl, createdAt: time.Now()}
	return true, nil
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}


