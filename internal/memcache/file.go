package memcache

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	errOpenFile   = errors.New("error open file")
	errCreateFile = errors.New("error create file")
	errDecodeJSON = errors.New("error decoding json")
)

// SaveFile
func (c *Cache) SaveFile(relativePath string) error {
	file, err := os.Create(relativePath)
	if err != nil {
		return errCreateFile
	}

	encoder := json.NewEncoder(file)

	c.mutex.RLock()

	defer c.mutex.RUnlock()
	defer file.Close()

	err = encoder.Encode(c.m)
	if err != nil {
		return err
	}

	return nil
}

// LoadFile
func (c *Cache) LoadFile(relativePath string) error {
	file, err := os.Open(relativePath)
	if err != nil {
		return errOpenFile
	}
	defer file.Close()

	newMap := make(map[string]interface{}, defaultSize)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.m)
	if err != nil {
		return errDecodeJSON
	}

	fmt.Println(newMap)

	return nil
}
