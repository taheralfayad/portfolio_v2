package v1

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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

func AddCoffeeCup(c *gin.Context, db *sql.DB) {
	var payload data.CoffeeCup

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	var brewDate time.Time
	err := db.QueryRow("SELECT brew_date FROM coffee WHERE id = $1", payload.CoffeeId).Scan(&brewDate)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	dateDrank, err := time.Parse(time.RFC3339, payload.DateDrank)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	daysAfterRoast := int16(dateDrank.Sub(brewDate).Hours() / 24)

	query := `
		INSERT INTO coffee_cup (coffee_id, temperature, date_drank, acidity, body, sweetness, water_type, grind_size, method, rating)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(
		query,
		payload.CoffeeId,
		payload.Temperature,
		payload.DateDrank,
		payload.Acidity,
		payload.Body,
		payload.Sweetness,
		payload.WaterType,
		payload.GrindSize,
		payload.Method,
		payload.Rating,
	)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	var id int
	err = db.QueryRow("SELECT lastval()").Scan(&id)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	response := data.CoffeeCup{
		ID:             id,
		CoffeeId:       payload.CoffeeId,
		Temperature:    payload.Temperature,
		DaysAfterRoast: daysAfterRoast,
		Acidity:        payload.Acidity,
		Body:           payload.Body,
		Sweetness:      payload.Sweetness,
		WaterType:      payload.WaterType,
		GrindSize:      payload.GrindSize,
		Method:         payload.Method,
		Rating:         payload.Rating,
	}

	messages.StatusCreated(c, response)
}

func GetCoffeeCups(c *gin.Context, db *sql.DB) {
	coffeeID := c.Query("coffee_id")
	if coffeeID == "" {
		messages.BadRequest(c, fmt.Errorf("coffee_id is required"))
		return
	}

	query := `
		SELECT 
			cc.id,
			cc.coffee_id,
			cc.temperature,
			cc.acidity,
			cc.body,
			cc.sweetness,
			cc.water_type,
			cc.grind_size,
			cc.method,
			cc.rating,
			c.brew_date,
			cc.date_drank
		FROM coffee_cup cc
		JOIN coffee c ON cc.coffee_id = c.id
		WHERE cc.coffee_id = $1
		ORDER BY date_drank DESC
	`

	rows, err := db.Query(query, coffeeID)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}
	defer rows.Close()

	var cups []data.CoffeeCup

	for rows.Next() {
		var cup data.CoffeeCup
		var brewDate time.Time
		var dateDrank time.Time

		err := rows.Scan(
			&cup.ID,
			&cup.CoffeeId,
			&cup.Temperature,
			&cup.Acidity,
			&cup.Body,
			&cup.Sweetness,
			&cup.WaterType,
			&cup.GrindSize,
			&cup.Method,
			&cup.Rating,
			&brewDate,
			&dateDrank,
		)
		if err != nil {
			messages.BadRequest(c, err)
			return
		}

		cup.DaysAfterRoast = int16(dateDrank.Sub(brewDate).Hours() / 24)

		cup.DateDrank = dateDrank.Format(time.RFC3339)

		cups = append(cups, cup)
	}

	messages.StatusCreated(c, cups)
}
