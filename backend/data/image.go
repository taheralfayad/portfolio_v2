package data

type ImageBase struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Site    string `json:"site"`
}

type ImagePayload struct {
	ImageBase
	Image string `json:"image"`
}

type ImageResponse struct {
	ImageBase
	ImageLink string `json:"image"`
}
