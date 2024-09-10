package entities

import "time"

type ShortURL struct {
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type CreateShortURLRequest struct {
	OriginalURL string  `json:"original_url" validate:"required,url" example:"https://www.google.com/"`
	CustomURL   *string `json:"custom_url,omitempty" validate:"omitempty,alphanum,gte=5,lt=20" example:"shorturl"`
	Duration    *int64  `json:"duration,omitempty" validate:"omitempty,gt=1,lt=90" example:"3"`
}

type CreateShortURLResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	Expires  string `json:"expires"`
}

type ShortURLResponseSuccess struct {
	Data CreateShortURLResponse `json:"data"`
}

type ShortURLResponseError struct {
	Error string `json:"error"`
}
