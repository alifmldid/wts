package server

import (
	"user-service/config"
	"user-service/user"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine) {
	db := config.GetDBConnection()

	userRepo := user.NewUserRepository(db)
	userUSecase := user.NewUserUsecase(userRepo)
	userController := user.NewUserController(userUSecase)

	registerUserRoute(r, userController)
}