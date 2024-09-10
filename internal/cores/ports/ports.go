package ports

import (
	"github.com/savioruz/short/internal/cores/entities"
	"time"
)

type ShortURLRepository interface {
	CreateShortURL(originalURL string, shortCode *string, duration *time.Duration) (*entities.ShortURL, error)
	GetLongURL(shortCode string) (string, error)
}
