package memcache

import (
	"errors"
	"runtime"
	"sync"
	"time"
)

var (
	defaultSize int

	KeyNotFound = errors.New("key not found")
	CacheIsOut  = errors.New("cache is out of date")
)

type Cache struct {
	// если установлено значение меньше или равно 0 — время жизни кеша бессрочно
	defaultExpiration time.Duration // продолжительность жизни кеша по-умолчанию
	// При установленном значении меньше или равно 0 — очистка и удаление просроченного кеша не происходит
	cleanupInterval time.Duration // интервал, через который запускается механизм очистки кеша

	m     map[string]Value
	mutex sync.RWMutex
}

type Value struct {
	Value      interface{} `json:"value"`
	createdAt  time.Time   `json:"created_at"`
	expiration int64       `json:"expiration"` // Актуальность кэша
}

// NewCache Create a new cache container.
// it will start gc automatically.
func NewCache(size int, expiration, cleanupInterval time.Duration) *Cache {
	newMap := make(map[string]Value, size)

	defaultSize = size

	cache := &Cache{
		defaultExpiration: expiration,
		cleanupInterval:   cleanupInterval,
		m:                 newMap,
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return cache
}

func (c *Cache) StartGC() {
	go c.GC()
}

// Set cache by key with duration.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.m[key] = Value{
		Value:      value,
		expiration: expiration,
		createdAt:  time.Now(),
	}
}

// Get get cached value by key.
func (c *Cache) Get(key string) (value interface{}, err error) {
	c.mutex.RLock()

	defer c.mutex.RUnlock()

	val, ok := c.m[key]
	if !ok {
		return nil, KeyNotFound
	}

	if val.expiration > 0 {
		if time.Now().UnixNano() > val.expiration {
			return nil, CacheIsOut
		}
	}

	return val.Value, nil
}

// GetMulti gets caches from memory.
// if non-existed or expired, return nil.
func (c *Cache) GetMulti(names []string) []interface{} {
	var rc []interface{}
	for _, name := range names {
		v, _ := c.Get(name)
		rc = append(rc, v)
	}
	return rc
}

// Delete remove cache by key.
func (c *Cache) Delete(key string) error {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	if _, ok := c.m[key]; !ok {
		return KeyNotFound
	}

	delete(c.m, key)

	return nil
}

// Exist check if cached value exists or not.
func (c *Cache) IsExist(key string) bool {

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	_, ok := c.m[key]

	return ok
}

// Expire check if cached value expired or not.
// if cache expire == true, cache not expire == false.
func (c *Cache) Expire(key string) (bool, error) {

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	val, ok := c.m[key]
	if !ok {
		return false, KeyNotFound
	}

	if time.Now().UnixNano() > val.expiration && val.expiration > 0 {
		return true, nil
	}

	return false, nil
}

// FlushAll clear all cache.
func (c *Cache) FlushAll() {
	newMap := make(map[string]Value, defaultSize)

	c.m = newMap
	runtime.GC()
}

func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if len(c.m) == 0 {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

// expiredKeys returns a list of "expired" keys
func (c *Cache) expiredKeys() (keys []string) {

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	for k, i := range c.m {
		if time.Now().UnixNano() > i.expiration && i.expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems removes keys from the passed list, in our case "expired"
func (c *Cache) clearItems(keys []string) {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	for _, k := range keys {
		delete(c.m, k)
	}
}
