package repositories

import (
	"fmt"
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

func (s *DB) SetKey(key, value string) string {
	return fmt.Sprintf("%s:%s", key, value)
}
