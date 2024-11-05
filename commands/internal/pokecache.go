package internal

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
	mu sync.Mutex // thread safety
	intervalDur time.Duration // reaping timing
}

func NewCache(interval time.Duration) *Cache{
	cacheMap := make(map[string]cacheEntry)
	newCache := &Cache{cacheMap: cacheMap,
	intervalDur: interval,}
	go newCache.reapLoop()
	return newCache
}

func(c *Cache) reapLoop() {
	ticker := time.NewTicker(c.intervalDur)
	defer ticker.Stop()

	for {
		<-ticker.C
		cutoff := time.Now().Add(-c.intervalDur)
		c.mu.Lock()
		for k, v := range c.cacheMap {
			if v.createdAt.Before(cutoff) {
				delete(c.cacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}


func(c *Cache) Add(key string, val []byte) (error) {

	c.mu.Lock()
	_, exists := c.Get(key)
	if exists {
		return fmt.Errorf("key already exists in cache")
	}

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	defer c.mu.Unlock()
	return nil
}

func(c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, exists := c.cacheMap[key]
	if exists {
		return data.val, true
	}
	return nil, false
}

