package repositories

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/pkg/constants"
)

func (s *DB) CreatePaste(title, content string, pasteID *string, duration *int8) (*entities.Paste, error) {
	id := uuid.NewString()[:6]
	if pasteID != nil && *pasteID != "" {
		// Check if custom key already exists in shorten or paste
		for _, keyType := range []string{"shorten", "paste"} {
			exists, err := s.checkCustomKey(keyType, *pasteID)
			if err != nil {
				return nil, err
			}

			if exists {
				return nil, fmt.Errorf("custom URL already exists in %s", keyType)
			}
		}

		id = *pasteID
	}

	key := s.setKey("paste", id)
	now, expiresAt, expireDuration := s.calcExpiration(duration, constants.DefaultDuration)

	paste := &entities.Paste{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: now,
		ExpiresAt: expiresAt,
	}

	set := s.cache.Set(key, paste, expireDuration)
	if set != nil {
		return nil, errors.New("failed to set cache")
	}

	return paste, nil
}

func (s *DB) GetPaste(pasteID string) (*entities.Paste, error) {
	var paste entities.Paste
	key := s.setKey("paste", pasteID)

	err := s.cache.Get(key, &paste)
	if err != nil {
		return nil, errors.New("could not retrieve paste data from cache")
	}

	return &paste, nil
}
