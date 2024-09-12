package rest

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/savioruz/short/internal/cores/services"
	"time"
)

type PasteHandler struct {
	service   *services.PasteService
	validator *validator.Validate
}

func NewPasteHandler(service *services.PasteService) *PasteHandler {
	return &PasteHandler{
		service:   service,
		validator: validator.New(),
	}
}

// CreatePaste function is a handler to create a paste
// @Description Create a paste
// @Summary create a paste
// @Tags Paste
// @Accept json
// @Produce json
// @Param data body CreatePasteRequest true "Create Paste Request"
// @Success 200 {object} PasteResponseSuccess
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /paste [post]
func (h *PasteHandler) CreatePaste(c *fiber.Ctx) error {
	var req CreatePasteRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleError(c, fiber.StatusBadRequest, err)
	}

	if err := h.validator.Struct(req); err != nil {
		return HandleError(c, fiber.StatusBadRequest, err)
	}

	p, err := h.service.CreatePaste(req.Title, req.Content, req.ID, req.Duration)
	if err != nil {
		return HandleError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(PasteResponseSuccess{
		Data: createPasteResponse{
			ID:      p.ID,
			URL:     fmt.Sprintf("%s/%s", c.BaseURL(), p.ID),
			Title:   p.Title,
			Content: p.Content,
			Created: p.CreatedAt.Format(time.RFC3339),
			Expires: p.ExpiresAt.Format(time.RFC3339),
		},
	})
}
