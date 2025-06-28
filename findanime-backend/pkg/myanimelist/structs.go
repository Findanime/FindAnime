package myanimelist

type ApiResponse struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

type Item struct {
	ImageURL string `json:"image_url"`
}
