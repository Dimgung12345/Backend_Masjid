package cache

import (
    "sync"
    "time"
)

// cacheItem menyimpan data generic + waktu expired
type cacheItem struct {
    data       interface{}
    expiration int64
}

// CacheManager adalah in-memory cache keyed by string (misalnya clientID + tipe data)
type CacheManager struct {
    mu    sync.RWMutex
    store map[string]cacheItem
    ttl   time.Duration
}

// NewCacheManager membuat instance baru dengan TTL default
func NewCacheManager(ttl time.Duration) *CacheManager {
    return &CacheManager{
        store: make(map[string]cacheItem),
        ttl:   ttl,
    }
}

// Get mengambil data dari cache jika masih valid
func (c *CacheManager) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, found := c.store[key]
    if !found {
        return nil, false
    }
    if time.Now().Unix() > item.expiration {
        // expired → anggap miss
        return nil, false
    }
    return item.data, true
}

// Set menyimpan data ke cache dengan TTL
func (c *CacheManager) Set(key string, data interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.store[key] = cacheItem{
        data:       data,
        expiration: time.Now().Add(c.ttl).Unix(),
    }
}

// Invalidate menghapus data dari cache
func (c *CacheManager) Invalidate(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()

    delete(c.store, key)
}