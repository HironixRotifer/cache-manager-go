package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	var key = "hello"
	var expectedResult = "world"

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	cache.Set(key, expectedResult, 0)

	actualResult, err := cache.Get(key)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGet(t *testing.T) {
	var key = "hello"
	var expectedResult = "world"

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	cache.Set(key, expectedResult, 0)

	actualResult, err := cache.Get(key)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGetMulti(t *testing.T) {
	var keys = []string{"hello", "frog", "mario", "cart"}
	var expectedResult = []string{"world", "lyagushka", "cart", "Tax"}

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	for index, key := range keys {
		cache.Set(key, expectedResult[index], 0)
	}

	actualResult := cache.GetMulti(keys)

	for index, result := range actualResult {
		assert.Equal(t, expectedResult[index], result)
	}
}

func TestDelete(t *testing.T) {
	var key = "hello"

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	cache.Set(key, "world", 0)

	exist := cache.IsExist(key)
	assert.Equal(t, true, exist)

	err := cache.Delete(key)
	assert.NoError(t, err)

	exist = cache.IsExist(key)
	assert.Equal(t, false, exist)

}

func TestIsExist(t *testing.T) {
	var key = "hello"

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)

	exist := cache.IsExist(key)
	assert.Equal(t, false, exist)

	cache.Set(key, "world", 0)

	exist = cache.IsExist(key)
	assert.Equal(t, true, exist)
}

func TestExpire(t *testing.T) {
	var key = "hello"

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	cache.Set(key, "world", 1*time.Nanosecond)
	time.Sleep(1 * time.Nanosecond)

	expire, err := cache.Expire(key)
	assert.NoError(t, err)
	assert.Equal(t, true, expire)

	cache.Set(key, "world", 1*time.Second)

	expire, err = cache.Expire(key)
	assert.NoError(t, err)

	assert.Equal(t, false, expire)
}

func TestFlushAll(t *testing.T) {
	var keys = []string{"hello", "frog", "mario", "cart"}
	var expectedResult = []string{"world", "lyagushka", "cart", "Tax"}
	var lenOldMap int

	cache := NewCache(5, 5*time.Minute, 5*time.Minute)
	for index, key := range keys {
		cache.Set(key, expectedResult[index], 0)
	}

	lenOldMap = len(cache.m)

	cache.FlushAll()

	assert.NotEqual(t, lenOldMap, len(cache.m))
	assert.Equal(t, 0, len(cache.m))
}
