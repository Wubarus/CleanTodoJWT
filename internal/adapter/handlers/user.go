package handlers

import (
	"CTodo/internal/core/domain"
	"CTodo/internal/core/services"
	"CTodo/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(UserService services.UserService) *UserHandler {
	return &UserHandler{
		svc: UserService,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user domain.User

	err := ctx.BindJSON(&user)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input")
	}

	_, err = h.svc.Register(user.Username, user.Email, user.Password)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (h *UserHandler) ReadUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.svc.ReadUser(id)

	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	const oper = "internal.adapter.handlers.user.UpdateUser"

	claims, err := utils.ValidateToken(ctx.Request.Header.Get("Authorization"), os.Getenv("JWT_SECRET"))
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := claims.GetSubject()
	if err != nil {
		fmt.Printf("failed to get userId %v: %v", oper, err)
	}

	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.svc.UpdateUser(userId, user.Email, user.Username, user.Password)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	const oper = "internal.adapter.handlers.user.DeleteUser"

	claims, err := utils.ValidateToken(ctx.Request.Header.Get("Authorization"), os.Getenv("JWT_SECRET"))
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := claims.GetSubject()
	if err != nil {
		fmt.Printf("failed to get userId %v: %v", oper, err)
	}

	err = h.svc.DeleteUser(userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
