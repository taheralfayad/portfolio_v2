package v1

import (
	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func AddUser(c *gin.Context, db *sql.DB) {
	var user data.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO users (name, password)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err = db.QueryRow(
		query,
		user.Name,
		hashedPassword,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, user)
}
