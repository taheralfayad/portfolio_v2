package v1

import (
	data "github.com/taheralfayad/portfolio_v2/data"

	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func GetAllTables(c *gin.Context, db *sql.DB) {
		rows, err := db.Query(`
			SELECT table_name
			FROM information_schema.tables
			WHERE table_schema = 'public'
				AND table_type = 'BASE TABLE';
		`)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer rows.Close()

		var tables []data.Table

		for rows.Next() {
			var t data.Table

			err := rows.Scan(
				&t.TableName,
			)
			
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			tables = append(tables, t)
		}

		c.JSON(http.StatusOK, tables)
}
