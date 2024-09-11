package repositories

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/pkg/constants"
	"time"
)

func (s *DB) CreateShortURL(originalURL string, shortCode *string, duration *int8) (*entities.ShortURL, error) {
	var shorten, key string
	if shortCode == nil || *shortCode == "" {
		shorten = uuid.NewString()[:6]
	} else {
		var shortURL entities.ShortURL
		key = s.SetKey("shorten", *shortCode)
		err := s.cache.Get(key, &shortURL)
		if err == nil || shortURL.ShortCode != "" {
			return nil, errors.New("custom url already exists")
		}
		shorten = *shortCode
	}

	key = s.SetKey("shorten", shorten)
	now := time.Now()
	var expiresAt time.Time
	var expireDuration time.Duration
	if duration != nil {
		expiresAt = now.Add((time.Duration(*duration)) * constants.DefaultDuration)
		expireDuration = (time.Duration(*duration)) * constants.DefaultDuration
	} else {
		expiresAt = now.Add(constants.DefaultDuration)
		expireDuration = constants.DefaultDuration
	}

	shortURL := &entities.ShortURL{
		OriginalURL: originalURL,
		ShortCode:   shorten,
		CreatedAt:   now,
		ExpiresAt:   expiresAt,
	}

	set := s.cache.Set(key, shortURL, expireDuration)
	if set != nil {
		return nil, errors.New("failed to set cache")
	}

	return shortURL, nil
}

func (s *DB) GetLongURL(shortCode string) (string, error) {
	var shortURL entities.ShortURL
	key := fmt.Sprintf("shorten:%s", shortCode)

	err := s.cache.Get(key, &shortURL)
	if err != nil {
		return "", errors.New("could not retrieve URL from cache")
	}

	return shortURL.OriginalURL, nil
}
