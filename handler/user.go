package handler

import (
	"bwaproject/auth"
	"bwaproject/helper"
	"bwaproject/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (handler *userHandler) RegisterUser(context *gin.Context) {
	var input user.RegisterUserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMsg := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := handler.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMsg := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userToken, err := handler.authService.GenerateToken(newUser.ID)
	
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusOK, "error", nil)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newUser, userToken)

	response := helper.APIResponse("Account Has Been Registered", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)
}

func (handler *userHandler) Login(context *gin.Context) {
	var input user.LoginUserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMsg := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := handler.userService.Login(input)
	if err != nil {
		errorMsg := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userToken, err := handler.authService.GenerateToken(loggedInUser.ID)
	
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusOK, "error", nil)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, userToken)

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)
}

func(handler *userHandler) CheckEmailAvailability(context *gin.Context) {
	var input user.CheckEmailInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMsg := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := handler.userService.IsEmailAvailable(input)
	
	if err != nil {
		errorMsg := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusOK, "error", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available" : isEmailAvailable}
	
	metaMessage := "Email Has been registered"

	if isEmailAvailable {
		metaMessage = "Email Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	context.JSON(http.StatusOK, response)

}

func(handler *userHandler) UploadAvatar(context *gin.Context) {
	file, err := context.FormFile("avatar")
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMsg)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = context.SaveUploadedFile(file,path)
	
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMsg)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = handler.userService.SaveAvatar(userID, path)
	if err != nil {
		errorMsg := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMsg)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	context.JSON(http.StatusOK, response)

}