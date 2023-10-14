package main

import (
	"fmt"
	"log"
	"os"

	"github.com/celpung/clean-gin-architecture/configs"
	"github.com/celpung/clean-gin-architecture/infrastructure"
	"github.com/celpung/clean-gin-architecture/internal/handler"
	"github.com/celpung/clean-gin-architecture/internal/repository"
	"github.com/celpung/clean-gin-architecture/internal/usecase"
	"github.com/celpung/clean-gin-architecture/middlewares"
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

	// infrastructure
	passwordSrv := infrastructure.NewPasswordService()
	jwtSrv := infrastructure.NewJwtService()

	/*
		users
	*/
	userRepo := repository.NewUserRepository(configs.DB)
	userUseCase := usecase.NewUserUseCase(userRepo, passwordSrv, jwtSrv)
	userHandler := handler.NewUserHandler(*userUseCase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/", middlewares.JWTMiddleware(configs.Admin), userHandler.GetAllUser)
		userRoutes.GET("/:id", middlewares.JWTMiddleware(configs.User), userHandler.GetUserByID)
		userRoutes.PUT("/:id", middlewares.JWTMiddleware(configs.User), userHandler.UpdateUser)
		userRoutes.DELETE("/:id", middlewares.JWTMiddleware(configs.Admin), userHandler.DeleteUser)
		userRoutes.POST("/sign-in", userHandler.SignIn)
	}

	// Start the server
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
