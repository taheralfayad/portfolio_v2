package v1

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	data "github.com/taheralfayad/portfolio_v2/data"
	messages "github.com/taheralfayad/portfolio_v2/messages"
	"github.com/taheralfayad/portfolio_v2/utils"
)

func AddCoffee(c *gin.Context, db *sql.DB) {
	var payload data.CoffeeRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	id := uuid.New()
	imageID := fmt.Sprintf("%s_image", id.String())
	imageLink := "/coffee_images/" + imageID

	err := utils.SaveBase64ImageToDisk(
		payload.Image,
		os.Getenv("STATIC_DIR")+imageLink,
	)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	query := `
		INSERT INTO coffee (name, roast_level, roaster_name, origin_country, processing, varietal, image, brew_date, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = db.Exec(
		query,
		payload.Name,
		payload.RoastLevel,
		payload.RoasterName,
		payload.OriginCountry,
		payload.Processing,
		payload.Varietal,
		imageLink,
		payload.Date,
		payload.Description,
	)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	messages.StatusCreated(c, "Coffee successfully submitted. Good job!")
}

func GetCoffees(c *gin.Context, db *sql.DB) {
	limit := c.DefaultQuery("limit", "0")

	query := `
		SELECT
			id,
			name,
			roast_level,
			roaster_name,
			origin_country,
			processing,
			image,
			varietal,
			brew_date,
			description
		FROM coffee
	`

	if limit != "0" {
		query += "LIMIT" + " " + limit
	}

	rows, err := db.Query(query)
	if err != nil {
		messages.BadRequest(c, err)
	}
	defer rows.Close()

	var coffees []data.CoffeeResponse

	for rows.Next() {
		var co data.CoffeeResponse

		err := rows.Scan(
			&co.ID,
			&co.Name,
			&co.RoastLevel,
			&co.RoasterName,
			&co.OriginCountry,
			&co.Processing,
			&co.Image,
			&co.Varietal,
			&co.Date,
			&co.Description,
		)

		co.Image = os.Getenv("ASSETS_URL") + co.Image

		if err != nil {
			messages.BadRequest(c, err)
		}

		coffees = append(coffees, co)
	}

	messages.StatusCreated(c, coffees)
}
