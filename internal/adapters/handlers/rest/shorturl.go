package rest

type CreateShortURLRequest struct {
	OriginalURL string  `json:"original_url" validate:"required,url" example:"https://www.google.com/"`
	CustomURL   *string `json:"custom_url,omitempty" validate:"omitempty,alphanum,gte=5,lt=20" example:"shorturl"`
	Duration    *int8   `json:"duration,omitempty" validate:"omitempty,gte=1,lt=90" example:"3"`
}

type createShortURLResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	Expires  string `json:"expires"`
}

type ShortURLResponseSuccess struct {
	Data createShortURLResponse `json:"data"`
}
