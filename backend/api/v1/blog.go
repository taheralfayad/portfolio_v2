package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"

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
	includeStr := c.DefaultQuery("include", "[]")

	var includeArr []string

	err := json.Unmarshal([]byte(includeStr), &includeArr)

	if err != nil {
		slog.Error("error while retrieving blogs", err)
		messages.InternalError(c, errors.New("include query param must be of type arr"))
		return
	}

	columns := []string{"id", "created_at", "metadata"}

	acceptableIncludeColumns := map[string]bool{"content": true}

	if len(includeArr) > len(acceptableIncludeColumns) {
		slog.Error("error while retrieving blogs", err)
		messages.BadRequest(c, errors.New("invalid include query param"))
		return
	}

	for _, inc := range includeArr {
		col, _ := acceptableIncludeColumns[inc]
		if !col {
			messages.InternalError(c, fmt.Errorf("unsupported include value: %s", inc))
			return
		}
		columns = append(columns, inc)
	}

	var response []data.Blog

	query := fmt.Sprintf(`
		SELECT %s
		FROM blog
		ORDER BY created_at DESC;
	`, strings.Join(columns, ", "))

	rows, err := db.Query(query)

	if err != nil {
		slog.Error("error while retrieving blogs", err)
		messages.InternalError(c, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var blog data.Blog

		if slices.Contains(includeArr, "content") {
			rows.Scan(
				&blog.ID,
				&blog.CreatedAt,
				&blog.Metadata,
				&blog.Content,
			)
		} else {
			rows.Scan(
				&blog.ID,
				&blog.CreatedAt,
				&blog.Metadata,
			)
		}

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

func EditBlog(c *gin.Context, db *sql.DB) {
	var payload data.Blog

	blogId := c.Param("id")

	blogIdInt, err := strconv.Atoi(blogId)

	if err != nil {
		messages.BadRequest(c, errors.New("id must be of type int"))
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	query := `
		UPDATE blog
		SET content = $2, metadata = $3
		WHERE id = $1
	`

	result, err := db.Exec(
		query,
		blogIdInt,
		payload.Content,
		payload.Metadata,
	)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	if rowsAffected == 0 {
		messages.NotFound(c, errors.New("blog not found"))
		return
	}

	messages.StatusCreated(c, "Blog updated, Good job!")
}
