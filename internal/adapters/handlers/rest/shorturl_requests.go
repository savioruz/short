package rest

type CreateShortURLRequest struct {
	OriginalURL string  `json:"original_url" validate:"required,url" example:"https://www.google.com/"`
	CustomURL   *string `json:"custom_url,omitempty" validate:"omitempty,alphanum,gte=5,lt=20" example:"shorturl"`
	Duration    *int64  `json:"duration,omitempty" validate:"omitempty,gt=1,lt=90" example:"3"`
}
