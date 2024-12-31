package cache 

type Cache struct {
	storage map[string]interface{}
}

func New() Cache {
	return Cache{
		storage: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

func (c Cache) Get(key string) interface{} {
	value := c.storage[key]
	return value
}

func (c Cache) Delete(key string) {
	delete(c.storage, key)
}


