package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	// map of cache entries
	cache map[string]cacheEntry
	// used for preventing race conditions
	mux *sync.Mutex
}

type cacheEntry struct {
	// actual cache value
	val []byte
	// used for clearing cache once expired
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	// anther go routine to clear the cache on set intervals
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// lock/unlock mux
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// lock/unlock mux
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// returns a tick after each interval, runs forever
	// no mux usage here as it will forever lock the resource
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// clear all cache entries after every interval
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	// lock/unlock mux
	c.mux.Lock()
	defer c.mux.Unlock()
	// set the threshold time to clear all entries
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		// delete entry if after threshold time
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
