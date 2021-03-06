package main

import (
	"bwaproject/auth"
	"bwaproject/handler"
	"bwaproject/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "newuser:newuser@tcp(127.0.0.1:3306)/golangbwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
		api.POST("/sessions", userHandler.Login)
		api.POST("/email_checkers",userHandler.CheckEmailAvailability)
		api.POST("/avatars",userHandler.UploadAvatar)
	}

	err = router.Run(":8080")

	if err != nil {
		fmt.Println("Error while running server", err.Error())
	}

}
