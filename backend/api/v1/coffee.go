package v1

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	data "github.com/taheralfayad/portfolio_v2/data"
	messages "github.com/taheralfayad/portfolio_v2/messages"
	utils "github.com/taheralfayad/portfolio_v2/utils"
)

func AddCoffeeRoast(c *gin.Context, db *sql.DB) {
	var payload data.Roast

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
		INSERT INTO
		roast
		(
			roast_level,
			roaster_name,
			roast_date,
			image,
			coffee_id
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = db.Exec(
		query,
		payload.RoastLevel,
		payload.RoasterName,
		payload.RoastDate,
		imageLink,
		payload.CoffeeID,
	)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	messages.StatusCreated(c, "Roast successfully submitted. Good job!")
}

func AddCoffee(c *gin.Context, db *sql.DB) {
	var payload data.Coffee

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	query := `
		INSERT INTO
			coffee
		(
			name,
			origin_country,
			processing,
			varietal,
			description
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.Exec(
		query,
		payload.Name,
		payload.OriginCountry,
		payload.Processing,
		payload.Varietal,
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
	includeRoasts := c.DefaultQuery("include_roasts", "false")

	var (
		limitInt int
		err      error
		limitVal interface{}
	)

	query := `
		SELECT id, name, origin_country, processing, varietal, description
		FROM (
				SELECT DISTINCT ON (c.id)
						c.id,
						c.name,
						c.origin_country,
						c.processing,
						c.varietal,
						c.description,
						r.roast_date
				FROM coffee AS c
				LEFT JOIN roast AS r ON r.coffee_id = c.id
				ORDER BY c.id, r.roast_date DESC
		) AS latest
		ORDER BY roast_date DESC
		LIMIT $1
	`

	if limit != "0" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			messages.BadRequest(c, err)
			return
		}
		limitVal = limitInt
	} else {
		limitVal = nil
	}

	rows, err := db.Query(query, limitVal)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}
	defer rows.Close()

	var coffees []data.Coffee

	for rows.Next() {
		var co data.Coffee

		err := rows.Scan(
			&co.ID,
			&co.Name,
			&co.OriginCountry,
			&co.Processing,
			&co.Varietal,
			&co.Description,
		)
		if err != nil {
			messages.BadRequest(c, err)
		}

		coffees = append(coffees, co)
	}

	if includeRoasts == "true" {
		for i := range coffees {
			coffeeID := coffees[i].ID
			var roasts []data.Roast

			query := `
				SELECT
					id,
					roast_level,
					roaster_name,
					roast_date,
					image,
					coffee_id
				FROM roast
				WHERE coffee_id = $1
			`

			rows, err := db.Query(query, coffeeID)
			if err != nil {
				messages.InternalError(c, err)
			}

			for rows.Next() {
				var roast data.Roast
				var imageID string

				err := rows.Scan(
					&roast.ID,
					&roast.RoastLevel,
					&roast.RoasterName,
					&roast.RoastDate,
					&imageID,
					&roast.CoffeeID,
				)
				if err != nil {
					messages.InternalError(c, err)
				}

				fmt.Println(roast)

				imageLink := os.Getenv("ASSETS_URL") + imageID

				roast.Image = imageLink

				roasts = append(roasts, roast)
			}

			coffees[i].Roasts = roasts
		}
	}

	messages.StatusCreated(c, coffees)
}

func AddCoffeeCup(c *gin.Context, db *sql.DB) {
	var payload data.CoffeeCup

	if err := c.ShouldBindJSON(&payload); err != nil {
		messages.BadRequest(c, err)
		return
	}

	var roastDate time.Time

	query := `
		SELECT 
			roast_date
		FROM roast
		WHERE id = $1
	`

	err := db.QueryRow(
		query,
		payload.RoastID,
	).Scan(&roastDate)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	dateDrank, err := time.Parse(time.RFC3339, payload.DateDrank)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}

	daysAfterRoast := int16(dateDrank.Sub(roastDate).Hours() / 24)

	query = `
		INSERT INTO
			coffee_cup
		(
			roast_id,
			temperature,
			date_drank,
			acidity,
			body,
			sweetness,
			water_type,
			grind_size,
			method,
			rating
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(
		query,
		payload.RoastID,
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
		RoastID:        payload.RoastID,
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
	roastID := c.Query("roast_id")
	if roastID == "" {
		messages.BadRequest(c, fmt.Errorf("roast_id is required"))
		return
	}

	query := `
		SELECT 
			cc.id,
			cc.roast_id,
			cc.temperature,
			cc.acidity,
			cc.body,
			cc.sweetness,
			cc.water_type,
			cc.grind_size,
			cc.method,
			cc.rating,
			r.roast_date,
			cc.date_drank
		FROM coffee_cup cc
		JOIN roast r ON cc.roast_id = r.id
		WHERE cc.roast_id = $1
		ORDER BY date_drank DESC
	`

	rows, err := db.Query(query, roastID)
	if err != nil {
		messages.BadRequest(c, err)
		return
	}
	defer rows.Close()

	var cups []data.CoffeeCup

	for rows.Next() {
		var cup data.CoffeeCup
		var roastDate time.Time
		var dateDrank time.Time

		err := rows.Scan(
			&cup.ID,
			&cup.RoastID,
			&cup.Temperature,
			&cup.Acidity,
			&cup.Body,
			&cup.Sweetness,
			&cup.WaterType,
			&cup.GrindSize,
			&cup.Method,
			&cup.Rating,
			&roastDate,
			&dateDrank,
		)
		if err != nil {
			messages.BadRequest(c, err)
			return
		}

		cup.DaysAfterRoast = int16(dateDrank.Sub(roastDate).Hours() / 24)

		cup.DateDrank = dateDrank.Format(time.RFC3339)

		cups = append(cups, cup)
	}

	messages.StatusCreated(c, cups)
}

func DuckDBify(c *gin.Context, db *sql.DB) {
	coffeeRows, err := db.Query(`
		SELECT
			name,
			origin_country,
			processing,
			varietal,
			description
		FROM coffee
	`)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}
	defer coffeeRows.Close()

	var coffees []data.CoffeeParquet
	for coffeeRows.Next() {
		var coffee data.CoffeeParquet
		if err := coffeeRows.Scan(
			&coffee.Name,
			&coffee.OriginCountry,
			&coffee.Processing,
			&coffee.Varietal,
			&coffee.Description,
		); err != nil {
			slog.Error(err.Error(), "error", err)
			messages.InternalError(c, errors.New("Something went wrong"))
			return
		}
		coffees = append(coffees, coffee)
	}

	roastRows, err := db.Query(`
		SELECT
			roast_level,
			roaster_name,
			roast_date,
			(SELECT name FROM coffee c WHERE c.id = coffee_id) AS coffee_name
		FROM roast
	`)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}
	defer roastRows.Close()

	var roasts []data.RoastParquet
	for roastRows.Next() {
		var roast data.RoastParquet
		if err := roastRows.Scan(
			&roast.RoastLevel,
			&roast.RoasterName,
			&roast.RoastDate,
			&roast.CoffeeName,
		); err != nil {
			slog.Error(err.Error(), "error", err)
			messages.InternalError(c, errors.New("Something went wrong"))
			return
		}
		roasts = append(roasts, roast)
	}

	coffeeCupRows, err := db.Query(`
		SELECT
			temperature,
			date_drank,
    	EXTRACT(DAY FROM (date_drank - (SELECT roast_date FROM roast r WHERE r.id = roast_id)))::int AS days_after_roast,
			acidity,
			body,
			sweetness,
			water_type,
			grind_size,
			method,
			rating,
			(SELECT roast_date FROM roast r WHERE r.id = roast_id) AS roast_name
		FROM coffee_cup
	`)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}
	defer coffeeCupRows.Close()

	var coffeeCups []data.CoffeeCupParquet
	for coffeeCupRows.Next() {
		var coffeeCup data.CoffeeCupParquet
		if err := coffeeCupRows.Scan(
			&coffeeCup.Temperature,
			&coffeeCup.DateDrank,
			&coffeeCup.DaysAfterRoast,
			&coffeeCup.Acidity,
			&coffeeCup.Body,
			&coffeeCup.Sweetness,
			&coffeeCup.WaterType,
			&coffeeCup.GrindSize,
			&coffeeCup.Method,
			&coffeeCup.Rating,
			&coffeeCup.RoastName,
		); err != nil {
			slog.Error(err.Error(), "error", err)
			messages.InternalError(c, errors.New("Something went wrong"))
			return
		}
		coffeeCups = append(coffeeCups, coffeeCup)
	}

	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)

	if err := utils.WriteParquet(zw, "coffee.parquet", coffees); err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}
	if err := utils.WriteParquet(zw, "roast.parquet", roasts); err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}
	if err := utils.WriteParquet(zw, "coffee_cup.parquet", coffeeCups); err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}

	if err := zw.Close(); err != nil {
		slog.Error(err.Error(), "error", err)
		messages.InternalError(c, errors.New("Something went wrong"))
		return
	}

	c.Header("Content-Disposition", `attachment; filename="export.zip"`)
	c.Data(http.StatusOK, "application/zip", buf.Bytes())
}
