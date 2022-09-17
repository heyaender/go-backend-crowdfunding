package handlers

import (
	"go-backend-crowdfunding/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get Users"})
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, newUser)
}
