package repositories

import (
	"github.com/savioruz/short/internal/adapters/cache"
	"github.com/savioruz/short/internal/cores/ports"
)

type DB struct {
	cache ports.CacheRepository
}

func NewDB(c *cache.RedisCache) *DB {
	return &DB{
		cache: c,
	}
}
