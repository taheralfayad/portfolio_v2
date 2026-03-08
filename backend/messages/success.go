package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusCreated(c *gin.Context, response any) {
	c.JSON(http.StatusCreated, response)
}

func StatusOk(c *gin.Context, response any) {
	c.JSON(http.StatusOK, response)
}
