package messages

import "github.com/gin-gonic/gin"

func BadRequest(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}

func Unauthorized(c *gin.Context, err error) {
	c.JSON(401, gin.H{"error": err.Error()})
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(403, gin.H{"error": err.Error()})
}

func NotFound(c *gin.Context, err error) {
	c.JSON(404, gin.H{"error": err.Error()})
}

func InternalError(c *gin.Context, err error) {
	c.JSON(500, gin.H{"error": err.Error()})
}
