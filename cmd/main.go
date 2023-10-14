package main

import (
	"log"

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

	// Connect to the database
	configs.ConnectDatabase()
	configs.AutoMigrage()

	userRepo := repository.NewUserRepository(configs.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/", userHandler.GetAllUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
	gin.SetMode(gin.DebugMode)

	// Start the server
	r.Run(":8080")
}
