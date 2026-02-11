package v1

import (
	"os"
	"fmt"
	"context"

	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"database/sql"
)

func AddImage(c *gin.Context, db *sql.DB, ctx context.Context, client *s3.Client) {
	var payload data.ImagePayload
	var response data.ImageResponse
	var err error

	if err = c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.New()
	imageID := fmt.Sprintf("%s_image", id.String())

	imageLink := fmt.Sprintf(
		"%s/%s/%s",
		os.Getenv("RUSTFS_PUBLIC_URL"),
		"portfolio-assets",
		imageID,
	)

	err = utils.UploadBase64Image(
		ctx,
		client,
		"portfolio-assets",
		imageID,
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
		INSERT INTO images (title, caption, image_link)
		VALUES ($1, $2, $3)
	`

	_, err = db.Exec(
		query,
		payload.Title,
		payload.Caption,
		imageLink,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response.Title = payload.Title
	response.Caption = payload.Caption
	response.ImageLink = imageLink

	c.JSON(http.StatusCreated, response)
}

func GetImages(c *gin.Context, db *sql.DB) {
	limit := c.DefaultQuery("limit", "0")

	query := `
		SELECT 
			id,
			title,
			caption,
			image_link
		FROM images
	`

	if limit != "0" {
		query += "LIMIT" + " " + limit
	}
	
	rows, err := db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var images []data.ImageResponse

	for rows.Next(){
		var i data.ImageResponse

		err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Caption,
			&i.ImageLink,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		images = append(images, i)
	}

	c.JSON(http.StatusOK, images)
}

func EditImage(c *gin.Context, db *sql.DB, ctx context.Context, client *s3.Client) {
	var i data.ImagePayload
	var err error

	if err = c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var imageS3Key string

	if i.Image != "" {
		var previousImageLink string

		query := (`
			SELECT
				image_link
			FROM images
			WHERE id = $1
		`)

		err := db.QueryRow(
			query,
			i.ID,
		).Scan(&previousImageLink)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		imageS3Key, err = utils.ExtractUUIDFromImageURL(previousImageLink)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = utils.UploadBase64Image(
			ctx,
			client,
			"portfolio-assets",
			imageS3Key,
			i.Image,
			"image/png",
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	query := `
		UPDATE images
			SET title = $1,
				caption = $2
		WHERE id = $3
	`

	_, err = db.Exec(
		query,
		i.Title,
		i.Caption,
		i.ID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response data.ImageResponse

	c.JSON(http.StatusCreated, response)
}
