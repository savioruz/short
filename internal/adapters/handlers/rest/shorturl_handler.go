package rest

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/savioruz/short/internal/cores/services"
	"time"
)

type ShortURLHandler struct {
	shortURLService *services.ShortURLService
	pasteService    *services.PasteService
	validator       *validator.Validate
}

func NewShortURLHandler(shortURLService *services.ShortURLService, pasteService *services.PasteService) *ShortURLHandler {
	return &ShortURLHandler{
		shortURLService: shortURLService,
		pasteService:    pasteService,
		validator:       validator.New(),
	}
}

// CreateShortURL function is a handler to shorten a URL
// @Description Shorten a URL
// @Summary shorten a URL
// @Tags ShortURL
// @Accept json
// @Produce json
// @Param data body CreateShortURLRequest true "Create Short URL Request"
// @Success 200 {object} ShortURLResponseSuccess
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shorten [post]
func (h *ShortURLHandler) CreateShortURL(c *fiber.Ctx) error {
	var req CreateShortURLRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleError(c, fiber.StatusBadRequest, errors.New("invalid request"))
	}

	if err := h.validator.Struct(req); err != nil {
		return HandleError(c, fiber.StatusBadRequest, err)
	}

	s, err := h.shortURLService.CreateShortURL(req.OriginalURL, req.CustomURL, req.Duration)
	if err != nil {
		return HandleError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(ShortURLResponseSuccess{
		Data: createShortURLResponse{
			URL:      s.OriginalURL,
			ShortURL: fmt.Sprintf("%s/%s", c.BaseURL(), s.ShortCode),
			Expires:  s.ExpiresAt.Format(time.RFC3339),
		},
	})
}

// ResolveURL function is a handler to resolve a short URL or paste data automatically
func (h *ShortURLHandler) ResolveURL(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return HandleError(c, fiber.StatusBadRequest, errors.New("invalid short code"))
	}

	// Check if the code corresponds to a short URL
	shortUrl, err := h.shortURLService.GetLongURL(code)
	if err == nil {
		return c.Redirect(shortUrl, fiber.StatusMovedPermanently)
	} else if errors.Is(err, redis.Nil) {
		return HandleError(c, fiber.StatusNotFound, errors.New("shorten URL not found"))
	}

	// If no short URL is found, check if it's a paste
	paste, err := h.pasteService.GetPaste(code)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(PasteResponseSuccess{
			Data: createPasteResponse{
				ID:      paste.ID,
				URL:     fmt.Sprintf("%s/%s", c.BaseURL(), paste.ID),
				Title:   paste.Title,
				Content: paste.Content,
				Created: paste.CreatedAt.Format(time.RFC3339),
				Expires: paste.ExpiresAt.Format(time.RFC3339),
			},
		})
	} else if errors.Is(err, redis.Nil) {
		return HandleError(c, fiber.StatusNotFound, errors.New("no matching short URL or paste found"))
	}

	return HandleError(c, fiber.StatusInternalServerError, err)
}
