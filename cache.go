package cache 

import (
	"errors"
	"sync"
	"time"
)

type ItemCache struct {
	value interface{}
	expiredAt time.Time
}

type Cache struct {
	storage sync.Map
}

func New() *Cache {
	storage := &Cache{
		storage: sync.Map{},
	}
	go storage.backgroundCacheCleaner()

	return storage
}

func (c *Cache) Get(key string) (interface{}, error) {
	load, ok := c.storage.Load(key)
	if !ok {
		return false, errors.New("value for this key not exists")
	}
	value, ok := load.(ItemCache)
	if !ok {
		return false, errors.New("type assertion error")
	}
	return value.value, nil
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) (bool, error) {
	_, ok := c.storage.Load(key)
	if ok {
		return false, errors.New("value for this key already exists")
	}
	c.storage.Store(key, ItemCache{value: value, expiredAt: time.Now().Add(ttl)})
	return true, nil
}

func (c *Cache) Delete(key string) {
	c.storage.Delete(key)
}

func (c *Cache) backgroundCacheCleaner() {
	timer := time.NewTicker(time.Second * 1)
	defer timer.Stop()

	for {
		<- timer.C
		c.storage.Range(func(key, v interface{}) bool {
			value, ok := v.(ItemCache)
			if !ok {
				return false
			}
			if time.Now().After(value.expiredAt) {
				c.storage.Delete(key)
				return true
			}
			return true
		})
	}

}

