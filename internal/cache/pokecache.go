package pokecache

import (
	"fmt"
	"sync"
	"time"
)



type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry // caching
	mu *sync.Mutex // thread safety
}

func NewCache(interval time.Duration) Cache{
	cacheMap := make(map[string]cacheEntry)
	newCache := Cache{cacheMap: cacheMap,
	mu: &sync.Mutex{},}
	go newCache.reapLoop(interval)
	return newCache
}

func(c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheMap {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, k)
		}
	}
}


func(c *Cache) Add(key string, val []byte) () {
	fmt.Println("ADDING KEY")
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func(c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, exists := c.cacheMap[key]
	return data.val, exists
}

