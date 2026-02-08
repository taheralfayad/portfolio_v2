package v1

import (
	data "github.com/taheralfayad/portfolio_v2/data"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func AddUser(c *gin.Context, db *sql.DB) {
	var user data.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO users (name, password)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err = db.QueryRow(
		query,
		user.Name,
		hashedPassword,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, name, password FROM users;
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var users []data.User

	for rows.Next() {
		var u data.User

		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Password,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}

func EditUser(c *gin.Context, db *sql.DB) {
	var u data.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	query := `
		UPDATE users
		SET name = $1,
			password = $2
		WHERE id = $3
	`

	_, err = db.Exec(
		query,
		u.Name,
		hashedPassword,
		u.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func Login(c *gin.Context, db *sql.DB) {
	var req data.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	var (
		userID       int
		passwordHash string
	)

	err := db.QueryRow(
		`SELECT id, password FROM users WHERE name = $1`,
		req.Name,
	).Scan(&userID, &passwordHash)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "database error",
		})
		return
	}

	if err := utils.CheckPasswordHash(
		req.Password, passwordHash,
	); err == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJWT(userID, req.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	c.SetCookie(
		"access_token",
		token,
		3600,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}
