package routes

import (
	"go-backend-crowdfunding/apps/controllers/users"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Context) {
	router := gin.Default()
	router.GET("/users", users.UsersHandler)
	router.Run()
}
