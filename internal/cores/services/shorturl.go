package services

import (
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/internal/cores/ports"
	"time"
)

type ShortURLService struct {
	repo ports.ShortURLRepository
}

func NewShortURLService(repo ports.ShortURLRepository) *ShortURLService {
	return &ShortURLService{
		repo: repo,
	}
}

func (s *ShortURLService) CreateShortURL(originalURL string, shortCode *string, duration *time.Duration) (*entities.ShortURL, error) {
	return s.repo.CreateShortURL(originalURL, shortCode, duration)
}

func (s *ShortURLService) GetLongURL(shortCode string) (string, error) {
	return s.repo.GetLongURL(shortCode)
}
