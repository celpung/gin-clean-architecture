package handler

import (
	"net/http"

	"github.com/celpung/clean-gin-architecture/internal/entity"
	"github.com/celpung/clean-gin-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: uc}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Register failed",
		})
		return
	}

	if err := uh.UserUseCase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register success",
	})
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	user, err := uh.UserUseCase.GetUserByID(3)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh *UserHandler) GetAllUser(c *gin.Context) {
	user, err := uh.UserUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	user, err := uh.UserUseCase.GetUserByID(3)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.UserUseCase.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	if err := uh.UserUseCase.DeleteUser(3); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data deleted",
	})
}
