package handler

import (
	"bwaproject/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (handler *userHandler) RegisterUser(context *gin.Context) {
	var input user.RegisterUserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	user, err := handler.userService.RegisterUser(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	context.JSON(http.StatusOK, user)
}
