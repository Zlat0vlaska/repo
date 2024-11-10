package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildErrorResponse(c *gin.Context, status int, logMessage string) {
	c.JSON(status, gin.H{"error": logMessage})
}

func BuildSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
