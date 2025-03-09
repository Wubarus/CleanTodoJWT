package utils

import "github.com/gin-gonic/gin"

// Json message response
type errorResponse struct {
	Message string `json:"message"`
}

// Error response
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
