package v1

import (
	data "github.com/taheralfayad/portfolio_v2/data"

	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func AddWorkExperience(c *gin.Context, db *sql.DB) {
		var we data.WorkExperience

		if err := c.ShouldBindJSON(&we); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := `
			INSERT INTO work_experiences (title, workplace, description, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, created_at
		`

		err := db.QueryRow(
			query,
			we.Title,
			we.Workplace,
			we.Description,
			we.StartDate,
			we.EndDate,
		).Scan(&we.ID, &we.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, we)
}

func GetWorkExperiences(c *gin.Context, db *sql.DB) {
    limit := c.DefaultQuery("limit", "0")
    
    query := `
        SELECT id, title, workplace, description, start_date, end_date, created_at
        FROM work_experiences
        ORDER BY start_date ASC NULLS LAST`
    
    if limit != "0" {
        query += " LIMIT " + limit
    }
    
    rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer rows.Close()

		var experiences []data.WorkExperience

		for rows.Next() {
			var we data.WorkExperience
			var startDate, endDate sql.NullString

			err := rows.Scan(
				&we.ID,
				&we.Title,
				&we.Workplace,
				&we.Description,
				&startDate,
				&endDate,
				&we.CreatedAt,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			if startDate.Valid {
				we.StartDate = &startDate.String
			}
			if endDate.Valid {
				we.EndDate = &endDate.String
			}

			experiences = append(experiences, we)
		}

		c.JSON(http.StatusOK, experiences)
}

func EditWorkExperience(c *gin.Context, db *sql.DB){
		var we data.WorkExperience

		if err := c.ShouldBindJSON(&we); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := `
			UPDATE work_experiences
			SET title = $1,
					workplace = $2,
					description = $3,
					start_date = $4,
					end_date = $5
			WHERE id = $6
		`

		_, err := db.Exec(
			query,
			we.Title,
			we.Workplace,
			we.Description,
			we.StartDate,
			we.EndDate,
			we.ID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, we)
}
