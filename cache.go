package cache 

import (
	"errors"
)

type Cache struct {
	storage map[string]interface{}
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	value := c.storage[key]
	return value
}

func (c *Cache) Set(key string, value interface{}) (bool, error) {
	existsValue := c.storage[key]
	if existsValue != nil {
		return false, errors.New("value for this key already exists")
	}
	c.storage[key] = value
	return true, nil
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}


