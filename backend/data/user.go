package data

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
