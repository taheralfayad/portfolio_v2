package data

type ImagePayload struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	Image string `json:"image"`
}

type ImageResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Caption string `json:"caption"`
	ImageLink string `json:"image"`
}
