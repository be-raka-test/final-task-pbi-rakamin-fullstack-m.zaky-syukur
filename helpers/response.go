package helpers

import (
	"github.com/gin-gonic/gin"
)

// JSONResponse sends a JSON response with the specified status code and message.
func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// JSONError sends a JSON error response with the specified status code and error message.
func JSONError(c *gin.Context, statusCode int, err string) {
	JSONResponse(c, statusCode, err, nil)
}
