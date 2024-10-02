package memcache

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveFile(t *testing.T) {
	var relativePath = "../../test.json"
	var expectedResult = `{"hello": {"value": "world"}}`

	cache := NewCache(1, 1*time.Minute, 1*time.Minute)
	cache.Set("hello", "world", 5*time.Minute)

	err := cache.SaveFile(relativePath)
	assert.NoError(t, err)

	file, err := os.ReadFile(relativePath)
	assert.NoError(t, err)

	actualResult := string(file)

	assert.JSONEq(t, expectedResult, actualResult)
}

func TestLoadFile(t *testing.T) {
	var relativePath = "../../test.json"
	var expectedResult = "world"

	cache := NewCache(1, 1*time.Minute, 1*time.Minute)

	err := cache.LoadFile(relativePath)
	assert.NoError(t, err)

	actualResult, err := cache.Get("hello")
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, actualResult)
}
