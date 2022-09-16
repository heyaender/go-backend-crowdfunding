package users

import (
	"go-backend-crowdfunding/databases"
	"go-backend-crowdfunding/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
	var users []schemas.Users
	databases.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}
