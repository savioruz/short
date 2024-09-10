package handlers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/savioruz/short/internal/cores/entities"
	"github.com/savioruz/short/internal/cores/services"
	"time"
)

type ShortURLHandler struct {
	service   *services.ShortURLService
	validator *validator.Validate
}

func NewShortURLHandler(service *services.ShortURLService) *ShortURLHandler {
	return &ShortURLHandler{
		service:   service,
		validator: validator.New(),
	}
}

// CreateShortURL function is a handler to shorten a URL
// @Description Shorten a URL
// @Summary shorten a URL
// @Tags ShortURL
// @Accept json
// @Produce json
// @Param data body entities.CreateShortURLRequest true "Create Short URL Request"
// @Success 200 {object} entities.ShortURLResponseSuccess
// @Failure 400 {object} entities.ShortURLResponseError
// @Failure 500 {object} entities.ShortURLResponseError
// @Router /shorten [post]
func (h *ShortURLHandler) CreateShortURL(c *fiber.Ctx) error {
	var req entities.CreateShortURLRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ShortURLResponseError{
			Error: "Invalid request",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ShortURLResponseError{
			Error: "Validation failed : " + err.Error(),
		})
	}

	s, err := h.service.CreateShortURL(req.OriginalURL, req.CustomURL, (*time.Duration)(req.Duration))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ShortURLResponseError{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entities.ShortURLResponseSuccess{
		Data: entities.CreateShortURLResponse{
			URL:      s.OriginalURL,
			ShortURL: fmt.Sprintf("%s/%s", c.BaseURL(), s.ShortCode),
			Expires:  s.ExpiresAt.Format(time.RFC3339),
		},
	})
}

func (h *ShortURLHandler) ResolveURL(c *fiber.Ctx) error {
	shortCode := c.Params("url")
	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ShortURLResponseError{
			Error: "Invalid request",
		})
	}

	originalURL, err := h.service.GetLongURL(shortCode)
	if errors.Is(err, redis.Nil) {
		return c.Status(fiber.StatusNotFound).JSON(entities.ShortURLResponseError{
			Error: "URL not found",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ShortURLResponseError{
			Error: err.Error(),
		})
	}

	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}
