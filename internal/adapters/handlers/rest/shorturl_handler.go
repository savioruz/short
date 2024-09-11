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

	s, err := h.service.CreateShortURL(req.OriginalURL, req.CustomURL, req.Duration)
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

func (h *ShortURLHandler) ResolveURL(c *fiber.Ctx) error {
	shortCode := c.Params("url")
	if shortCode == "" {
		return HandleError(c, fiber.StatusBadRequest, errors.New("invalid short code"))
	}

	originalURL, err := h.service.GetLongURL(shortCode)
	if errors.Is(err, redis.Nil) {
		return HandleError(c, fiber.StatusNotFound, errors.New("url not found"))
	} else if err != nil {
		return HandleError(c, fiber.StatusInternalServerError, err)
	}

	return c.Redirect(originalURL, fiber.StatusMovedPermanently)
}
