package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/savioruz/short/internal/cores/entities"
	"time"
)

func (s *DB) CreateShortURL(originalURL string, shortCode *string, duration *int8) (*entities.ShortURL, error) {
	var customCode string
	if shortCode == nil || *shortCode == "" {
		customCode = uuid.NewString()[:6]
	} else {
		var shortURL entities.ShortURL
		err := s.cache.Get(*shortCode, &shortURL)
		if err == nil || shortURL.ShortCode != "" {
			return nil, errors.New("custom url already exists")
		}
		customCode = *shortCode
	}

	now := time.Now()
	defaultDuration := 24 * time.Hour
	var expiresAt time.Time
	var expireDuration time.Duration
	if duration != nil {
		expiresAt = now.Add((time.Duration(*duration)) * defaultDuration)
		expireDuration = (time.Duration(*duration)) * defaultDuration
	} else {
		expiresAt = now.Add(defaultDuration)
		expireDuration = defaultDuration
	}

	shortURL := &entities.ShortURL{
		OriginalURL: originalURL,
		ShortCode:   customCode,
		CreatedAt:   now,
		ExpiresAt:   expiresAt,
	}

	set := s.cache.Set(customCode, shortURL, expireDuration)
	if set != nil {
		return nil, errors.New("failed to set cache")
	}

	return shortURL, nil
}

func (s *DB) GetLongURL(shortCode string) (string, error) {
	var shortURL entities.ShortURL

	err := s.cache.Get(shortCode, &shortURL)
	if err != nil {
		return "", errors.New("could not retrieve URL from cache")
	}

	return shortURL.OriginalURL, nil
}
