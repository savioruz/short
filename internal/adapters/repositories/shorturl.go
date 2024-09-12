package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/pkg/constants"
)

func (s *DB) CreateShortURL(originalURL string, shortCode *string, duration *int8) (*entities.ShortURL, error) {
	shorten := uuid.NewString()[:6]
	if shortCode != nil && *shortCode != "" {
		exist, err := s.checkCustomKey("shorten", *shortCode)
		if err != nil {
			return nil, err
		}

		if exist {
			return nil, errors.New("shorten custom url already exists")
		}

		shorten = *shortCode
	}

	key := s.setKey("shorten", shorten)
	now, expiresAt, expireDuration := s.calcExpiration(duration, constants.DefaultDuration)

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
	key := s.setKey("shorten", shortCode)

	err := s.cache.Get(key, &shortURL)
	if err != nil {
		return "", errors.New("could not retrieve URL from cache")
	}

	return shortURL.OriginalURL, nil
}
