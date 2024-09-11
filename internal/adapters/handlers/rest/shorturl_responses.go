package rest

type createShortURLResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	Expires  string `json:"expires"`
}

type ShortURLResponseSuccess struct {
	Data createShortURLResponse `json:"data"`
}
