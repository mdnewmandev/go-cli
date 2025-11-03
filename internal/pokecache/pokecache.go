package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries 	 map[string]cacheEntry
	mutex   	 sync.Mutex
}

func NewCache(reapInterval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go cache.reapLoop(reapInterval)

	return cache
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	fmt.Printf("Cache entry added: %s\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(reapInterval time.Duration) {
	ticker := time.NewTicker(reapInterval)
	for range ticker.C {
		c.reapEntries(reapInterval)
	}
}

func (c *Cache) reapEntries(reapInterval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > reapInterval {
			delete(c.entries, key)
			fmt.Printf("Cache entry expired: %s\n", key)
		}
	}
}