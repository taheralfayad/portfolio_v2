package data

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GithubLink  string `json:"github_link"`
	BlogLink    string `json:"blog_link"`
	Type        string `json:"type"`
}

type ProjectPayload struct {
	Project
	Image string `json:"image"`
}

type ProjectResponse struct {
	Project
	ImageLink string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
