package routes

import (
	"go-backend-crowdfunding/databases"
	"go-backend-crowdfunding/handlers"
	"go-backend-crowdfunding/users"

	"github.com/gin-gonic/gin"
)

func LaunchApp() {

	// Creating a new user repository, service, and handler.
	userRepository := users.NewRepository(databases.DB)
	userService := users.NewService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize the router
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
	}
	router.Run()

}
