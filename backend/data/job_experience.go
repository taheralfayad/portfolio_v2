package data

type WorkExperience struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Workplace   string     `json:"workplace" binding:"required"`
	Description string     `json:"description"`
	StartDate   *string    `json:"start_date"`
	EndDate     *string    `json:"end_date"`
	CreatedAt   string     `json:"created_at"`
}
