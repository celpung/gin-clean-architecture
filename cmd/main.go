package main

import (
	"fmt"
	"log"
	"os"

	"github.com/celpung/clean-gin-architecture/configs"
	"github.com/celpung/clean-gin-architecture/internal/handler"
	"github.com/celpung/clean-gin-architecture/internal/repository"
	"github.com/celpung/clean-gin-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env", err)
	}
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Connect to the database
	configs.ConnectDatabase()
	configs.AutoMigrage()

	userRepo := repository.NewUserRepository(configs.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/", userHandler.GetAllUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Start the server
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
