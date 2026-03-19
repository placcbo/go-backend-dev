package main

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode   string `json:"shortcode"`
	OriginalURL string `json:"originalurl"`
	ShortURL    string `json:"short_url"`
}

type StatsResponse struct {
	ShortCode   string `json:"short"`
	OriginalURL string `json:"url"`
	Clicks      int    `json:"clicks"`
	CreatedAt   string `json:"created_at"`
}
