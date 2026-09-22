package v1

import (
	"database/sql"
	"errors"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	data "github.com/taheralfayad/portfolio_v2/data"
	messages "github.com/taheralfayad/portfolio_v2/messages"
)

func GetBlog(c *gin.Context, db *sql.DB) {
	var response data.Blog

	blogId := c.Param("id")

	blogIdInt, err := strconv.Atoi(blogId)

	if err != nil {
		slog.Error("an error occured", err)
		messages.BadRequest(c, err)
		return
	}

	query := `
		SELECT
			created_at,
			content,
			metadata
		FROM blog
		WHERE id = $1
	`

	err = db.QueryRow(query, blogIdInt).Scan(&response.CreatedAt, &response.Content, &response.Metadata)

	if err != nil {
		slog.Error("an error occured:", err)
		messages.BadRequest(c, errors.New("something went wrong"))
		return
	}

	messages.StatusOk(c, response)
}

func GetBlogs(c *gin.Context, db *sql.DB) {
	var response []data.Blog

	query := `
		SELECT
			id,
			created_at,
			metadata
		FROM blog
	`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		slog.Error("error while retrieving blogs", err)
		messages.InternalError(c, err)
	}

	for rows.Next() {
		var blog data.Blog

		rows.Scan(
			&blog.ID,
			&blog.CreatedAt,
			&blog.Metadata,
		)

		response = append(response, blog)

	}

	messages.StatusOk(c, response)
}

func AddBlog(c *gin.Context, db *sql.DB) {
	var payload data.Blog

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	query := `
		INSERT INTO blog
		(
			content,
			metadata
		)
		VALUES ($1, $2)
	`

	if _, err := db.Exec(
		query,
		payload.Content,
		payload.Metadata,
	); err != nil {
		messages.BadRequest(c, err)
		return
	}

	messages.StatusCreated(c, "Blog created. Good job!")
}
