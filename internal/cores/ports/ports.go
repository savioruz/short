package ports

import (
	"github.com/savioruz/short/internal/cores/entities"
)

type ShortURLRepository interface {
	CreateShortURL(originalURL string, shortCode *string, duration *int8) (*entities.ShortURL, error)
	GetLongURL(shortCode string) (string, error)
}
