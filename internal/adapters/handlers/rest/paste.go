package rest

type CreatePasteRequest struct {
	ID       *string `json:"paste_id,omitempty" validate:"omitempty,alphanum,gte=5,lt=20" example:"optional"`
	Title    string  `json:"title" validate:"required,gte=5,lt=50" example:"Title"`
	Content  string  `json:"content" validate:"required,gte=5" example:"Your content"`
	Duration *int8   `json:"duration,omitempty" validate:"omitempty,gte=1,lt=90" example:"3"`
}

type createPasteResponse struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

type PasteResponseSuccess struct {
	Data createPasteResponse `json:"data"`
}
