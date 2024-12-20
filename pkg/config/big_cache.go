package config

import (
	"time"

	"github.com/allegro/bigcache"
)

func NewBigCache() *bigcache.BigCache {
	config := bigcache.DefaultConfig(30 * time.Minute)
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
	}

	return cache
}
