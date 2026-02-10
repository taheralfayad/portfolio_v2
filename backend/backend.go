package main

import (
	v1 "github.com/taheralfayad/portfolio_v2/api/v1"
	utils "github.com/taheralfayad/portfolio_v2/utils"

	"os"
	"context"
	"log"
	"errors"
	"net/http"

	"fmt"
	"database/sql"
	
	"github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
	access_key_id := os.Getenv("RUSTFS_ACCESS_KEY_ID")
	secret_access_key := os.Getenv("RUSTFS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("RUSTFS_ENDPOINT_URL")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName,
	)

	cfg := aws.Config{
			EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
							URL: endpoint,
					}, nil
			}),
			Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(access_key_id, secret_access_key, "")),
	}


	ctx := context.Background()

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
	})

	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String("portfolio-assets"),
	})

	if err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr) {
					if apiErr.ErrorCode() != "BucketAlreadyOwnedByYou" {
						log.Fatalf("create bucket failed: %v", err)
					}
			}
	}

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
		v1.AddProject(c, db, ctx, client)
	})

	auth.POST("/users", func(c *gin.Context) {
		v1.AddUser(c, db)
	})

	auth.POST("/skills", func(c *gin.Context) {
		v1.AddSkill(c, db)
	})

	auth.POST("/images", func(c *gin.Context) {
		v1.AddImage(c, db, ctx, client)
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

	auth.PUT("/work-experiences", func(c *gin.Context) {
		v1.EditWorkExperience(c, db)
	})

	auth.PUT("/projects", func(c *gin.Context) {
		v1.EditProject(c, db, ctx, client)
	})

	auth.PUT("/users", func(c *gin.Context) {
		v1.EditUser(c, db)
	})

	auth.PUT("/skills", func(c *gin.Context) {
		v1.EditSkills(c, db)
	})

	auth.PUT("/images", func(c *gin.Context) {
		v1.EditImage(c, db, ctx, client)
	})

	r.Run()
}
