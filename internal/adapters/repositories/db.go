package repositories

import (
	"fmt"
	"github.com/savioruz/short/internal/adapters/cache"
	"github.com/savioruz/short/internal/cores/ports"
	"time"
)

type DB struct {
	cache ports.CacheRepository
}

func NewDB(c *cache.RedisCache) *DB {
	return &DB{
		cache: c,
	}
}

func (s *DB) setKey(key, value string) string {
	return fmt.Sprintf("%s:%s", key, value)
}

func (s *DB) checkCustomKey(keyType, customKey string) (bool, error) {
	var data interface{}
	key := s.setKey(keyType, customKey)

	err := s.cache.Get(key, &data)
	if err != nil {
		if err.Error() == "cache miss for key: "+key {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *DB) calcExpiration(duration *int8, defaultDuration time.Duration) (time.Time, time.Time, time.Duration) {
	now := time.Now()
	expireDuration := defaultDuration

	if duration != nil {
		expireDuration = time.Duration(*duration) * defaultDuration
	}

	expiresAt := now.Add(expireDuration)
	return now, expiresAt, expireDuration
}
