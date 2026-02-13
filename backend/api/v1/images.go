package v1

import (
	"os"
	"fmt"

	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"database/sql"
)

func AddImage(c *gin.Context, db *sql.DB) {
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
	imageLink := "/hero_images/" + imageID

	err = utils.SaveBase64ImageToDisk(
		payload.Image,
		os.Getenv("STATIC_DIR") + imageLink,
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

		i.ImageLink = os.Getenv("ASSETS_URL") + i.ImageLink

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

func EditImage(c *gin.Context, db *sql.DB) {
	var i data.ImagePayload
	var err error

	if err = c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if i.Image != "" {
		var imageLink string

		query := (`
			SELECT
				image_link
			FROM images
			WHERE id = $1
		`)

		err := db.QueryRow(
			query,
			i.ID,
		).Scan(&imageLink)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = utils.SaveBase64ImageToDisk(
			i.Image,
			os.Getenv("STATIC_DIR") + imageLink,
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
