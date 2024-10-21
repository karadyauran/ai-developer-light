package controller

import (
	"net/http"
	"wellness_api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userDto service.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.UserService.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userDto service.UserDTO
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.UserService.UpdateUser(id, userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uc.UserService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}