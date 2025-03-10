package routes

import (
	"github.com/gin-gonic/gin"
)

// UserHandler guarantee that its instance has listed methods
type UserHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

// NewUserRouter designed to manage user routes
func NewUserRouter(router *gin.Engine, userHandler UserHandler) *gin.RouterGroup {
	userRouter := router.Group("")
	{
		router.POST("/auth/register", userHandler.Register)
		router.POST("/auth/login", userHandler.Login)
		router.GET("/users/:id", userHandler.GetUser)
		router.PUT("/users", userHandler.UpdateUser)
		router.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return userRouter
}
