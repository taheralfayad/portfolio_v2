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

func GetSkills(c *gin.Context, db *sql.DB){
	rows, err := db.Query(
		`SELECT
			id, name, category, blog_link, created_at
		 FROM skills;
		`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	
	var skills []data.Skill

	for rows.Next() {
		var sk data.Skill

		err := rows.Scan(
			&sk.ID,
			&sk.Name,
			&sk.Category,
			&sk.BlogLink,
			&sk.CreatedAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		skills = append(skills, sk)
	}

	c.JSON(http.StatusOK, skills)
}

func EditSkills(c *gin.Context, db *sql.DB) {
	var sk data.Skill
	var err error

	if err = c.ShouldBindJSON(&sk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		UPDATE skills
		SET name = $1,
			category = $2,
			blog_link = $3
		WHERE id = $4
	`

	_, err = db.Exec(
		query,
		sk.Name,
		sk.Category,
		sk.BlogLink,
		sk.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, sk)
}
