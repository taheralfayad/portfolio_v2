package v1

import (
	data "github.com/taheralfayad/portfolio_v2/data"

	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func AddSkill(c *gin.Context, db *sql.DB) {
		var skill data.Skill

		if err := c.ShouldBindJSON(&skill); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := `
			INSERT INTO skills (name, category, blog_link)
			VALUES ($1, $2, $3)
			RETURNING id, created_at
		`

		err := db.QueryRow(
			query,
			skill.Name,
			skill.Category,
			skill.BlogLink,
		).Scan(&skill.ID, &skill.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, skill)
}
