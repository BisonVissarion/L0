package cacheMemory

import (
	"sync"

	"time"
)

type Cache struct {
	sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value      interface{}
	Expiration int64
	Created    time.Time
}

func New() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64
	c.Lock()
	defer c.Unlock()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}

	}
	return item.Value, true
}
