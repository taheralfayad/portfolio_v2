package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	v1 "github.com/taheralfayad/portfolio_v2/api/v1"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	auth := r.Group("/")
	auth.Use(AuthMiddleware())

	auth.POST("/work-experiences", func(c *gin.Context) {
		v1.AddWorkExperience(c, db)
	})

	auth.POST("/projects", func(c *gin.Context) {
		v1.AddProject(c, db)
	})

	auth.POST("/users", func(c *gin.Context) {
		v1.AddUser(c, db)
	})

	auth.POST("/skills", func(c *gin.Context) {
		v1.AddSkill(c, db)
	})

	auth.POST("/images", func(c *gin.Context) {
		v1.AddImage(c, db)
	})

	auth.POST("/coffees", func(c *gin.Context) {
		v1.AddCoffee(c, db)
	})

	auth.POST("/coffee-cups", func(c *gin.Context) {
		v1.AddCoffeeCup(c, db)
	})

	auth.POST("/roasts", func(c *gin.Context) {
		v1.AddCoffeeRoast(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		v1.Login(c, db)
	})

	r.GET("/all-tables", func(c *gin.Context) {
		v1.GetAllTables(c, db)
	})

	r.GET("/users", func(c *gin.Context) {
		v1.GetUsers(c, db)
	})

	r.GET("/work-experiences", func(c *gin.Context) {
		v1.GetWorkExperiences(c, db)
	})

	r.GET("/projects", func(c *gin.Context) {
		v1.GetProjects(c, db)
	})

	r.GET("/skills", func(c *gin.Context) {
		v1.GetSkills(c, db)
	})

	r.GET("/images", func(c *gin.Context) {
		v1.GetImages(c, db)
	})

	r.GET("/coffees", func(c *gin.Context) {
		v1.GetCoffees(c, db)
	})

	r.GET("/coffee-cups", func(c *gin.Context) {
		v1.GetCoffeeCups(c, db)
	})

	auth.PUT("/work-experiences", func(c *gin.Context) {
		v1.EditWorkExperience(c, db)
	})

	auth.PUT("/projects", func(c *gin.Context) {
		v1.EditProject(c, db)
	})

	auth.PUT("/users", func(c *gin.Context) {
		v1.EditUser(c, db)
	})

	auth.PUT("/skills", func(c *gin.Context) {
		v1.EditSkills(c, db)
	})

	auth.PUT("/images", func(c *gin.Context) {
		v1.EditImage(c, db)
	})

	r.Run()
}
