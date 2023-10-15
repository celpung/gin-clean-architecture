package routers

import (
	"github.com/celpung/clean-gin-architecture/configs"
	"github.com/celpung/clean-gin-architecture/infrastructure"
	"github.com/celpung/clean-gin-architecture/internal/handler"
	"github.com/celpung/clean-gin-architecture/internal/repository"
	"github.com/celpung/clean-gin-architecture/internal/usecase"
	"github.com/celpung/clean-gin-architecture/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	passwordSrv := infrastructure.NewPasswordService()
	jwtSrv := infrastructure.NewJwtService()

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
}
