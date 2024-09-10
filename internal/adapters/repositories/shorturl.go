package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/savioruz/short/internal/cores/entities"
	"time"
)

func (s *DB) CreateShortURL(originalURL string, shortCode *string, duration *time.Duration) (*entities.ShortURL, error) {
	var customCode string
	if shortCode == nil || *shortCode == "" {
		customCode = uuid.NewString()[:6]
	} else {
		existingURL := s.cache.Get(*shortCode, &entities.ShortURL{})
		if existingURL != nil {
			return nil, errors.New("custom url already exists")
		}
		customCode = *shortCode
	}

	now := time.Now()
	defaultDuration := 24 * time.Hour
	var expiresAt time.Time
	if duration != nil {
		expiresAt = now.Add(*duration * defaultDuration)
	} else {
		expiresAt = now.Add(defaultDuration)
	}

	shortURL := &entities.ShortURL{
		OriginalURL: originalURL,
		ShortCode:   customCode,
		CreatedAt:   now,
		ExpiresAt:   expiresAt,
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
