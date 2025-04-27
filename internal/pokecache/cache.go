package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheLock *sync.Mutex
	Entries   map[string]cacheEntry
	interval  time.Duration
}

func NewCache(Interval time.Duration) *Cache {
	cache := Cache{
		Entries:   make(map[string]cacheEntry),
		interval:  Interval,
		cacheLock: &sync.Mutex{},
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(url string, body []byte) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()
	c.Entries[url] = cacheEntry{
		createdAt: time.Now(),
		val:       body,
	}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()
	entry, ok := c.Entries[url]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		Now := time.Now()
		c.cacheLock.Lock()
		for k, v := range c.Entries {
			if Now.Sub(v.createdAt) > c.interval {
				delete(c.Entries, k)
			}
		}
		c.cacheLock.Unlock()
	}
}
