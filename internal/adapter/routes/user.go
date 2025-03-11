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
		userRouter.POST("/auth/register", userHandler.Register)
		userRouter.POST("/auth/login", userHandler.Login)
		userRouter.GET("/users/:id", userHandler.GetUser)
		userRouter.PUT("/users", userHandler.UpdateUser)
		userRouter.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return userRouter
}
