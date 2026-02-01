package v1

import (
	"os"
	"fmt"
	"context"

	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"database/sql"
)

func AddProject(c *gin.Context, db *sql.DB, ctx context.Context, client *s3.Client) {
		var payload data.ProjectPayload
		var response data.ProjectResponse
		var err error

		if err = c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = utils.UploadBase64Image(
			ctx,
			client,
			"portfolio-assets",
			fmt.Sprintf("%s_image", payload.Name),
			payload.Image,
			"image/png",
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		query := `
			INSERT INTO projects (name, description, github_link, image_link, blog_link, type)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at
		`

		imageLink := fmt.Sprintf(
			"%s/%s/%s",
			os.Getenv("RUSTFS_ENDPOINT_URL"),
			"portfolio-assets",
			fmt.Sprintf("%s_image", payload.Name),
		)

		err = db.QueryRow(
			query,
			payload.Name,
			payload.Description,
			payload.GithubLink,
			imageLink,
			payload.BlogLink,
			payload.Type,
		).Scan(&response.ID, &response.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		response.Name = payload.Name
		response.Description = payload.Description
		response.GithubLink = payload.GithubLink
		response.BlogLink = payload.BlogLink
		response.Type = payload.Type

		c.JSON(http.StatusCreated, response)

}
