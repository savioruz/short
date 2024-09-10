package repositories

import "github.com/savioruz/short/internal/adapters/cache"

type DB struct {
	cache *cache.RedisCache
}

func NewDB(c *cache.RedisCache) *DB {
	return &DB{
		cache: c,
	}
}
