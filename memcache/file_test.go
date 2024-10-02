package memcache

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveFile(t *testing.T) {
	var relativePath = "../../test.json"
	var expectedResult = "world"
	var actualResult = make(map[string]Value, 1)

	cache := NewCache(1, 1*time.Minute, 1*time.Minute)
	cache.Set("hello", "world", 5*time.Minute)

	err := cache.SaveFile(relativePath)
	assert.NoError(t, err)

	file, err := os.ReadFile(relativePath)
	assert.NoError(t, err)

	err = json.Unmarshal(file, &actualResult)
	assert.NoError(t, err)

	assert.Equal(t, expectedResult, actualResult["hello"].Value)
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
