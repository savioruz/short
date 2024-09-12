package services

import (
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/internal/cores/ports"
)

type PasteService struct {
	repo ports.PasteRepository
}

func NewPasteService(repo ports.PasteRepository) *PasteService {
	return &PasteService{
		repo: repo,
	}
}

func (s *PasteService) CreatePaste(title, content string, pasteID *string, duration *int8) (*entities.Paste, error) {
	return s.repo.CreatePaste(title, content, pasteID, duration)
}

func (s *PasteService) GetPaste(pasteID string) (*entities.Paste, error) {
	return s.repo.GetPaste(pasteID)
}
