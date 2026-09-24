package data

type Blog struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
	Metadata  string `json:"metadata"`
}
