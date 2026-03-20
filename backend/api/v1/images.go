package v1

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		os.Getenv("STATIC_DIR")+imageLink,
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

	var (
		rows *sql.Rows
		err  error
	)

	if limit != "0" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		rows, err = db.Query(`
			SELECT id, title, caption, image_link
			FROM images
			ORDER BY RANDOM()
			LIMIT $1
		`, limitInt)
	} else {
		rows, err = db.Query(`
			SELECT id, title, caption, image_link
			FROM images
			ORDER BY RANDOM()
		`)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var images []data.ImageResponse
	for rows.Next() {
		var i data.ImageResponse
		if err := rows.Scan(&i.ID, &i.Title, &i.Caption, &i.ImageLink); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		i.ImageLink = os.Getenv("ASSETS_URL") + i.ImageLink
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
			os.Getenv("STATIC_DIR")+imageLink,
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
