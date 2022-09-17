package routes

import (
	"go-backend-crowdfunding/databases"
	"go-backend-crowdfunding/handlers"
	"go-backend-crowdfunding/users"

	"github.com/gin-gonic/gin"
)

func LaunchApp() {

	// WebServer := fmt.Sprintf("Web Server is running on port %s", configs.WEB_PORT)
	// log.Println(WebServer)
	// log.Fatal(http.ListenAndServe(configs.WEB_PORT, nil))

	userRepository := users.NewRepository(databases.DB)
	userService := users.NewService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.RegisterUser)
	}
	router.Run()

}
