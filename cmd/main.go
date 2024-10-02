package main

import (
	"time"

	"github.com/HironixRotifer/go-memorycache-manager/internal/memcache"
)

func main() {
	c := memcache.NewCache(1, 1*time.Minute, 1*time.Minute)
	c.
}
