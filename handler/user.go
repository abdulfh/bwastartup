package handler

import (
	"bwaproject/helper"
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
		response := helper.APIResponse("Register Account Failed", http.StatusOK, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := handler.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusOK, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Account Has Been Registered", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)
}
