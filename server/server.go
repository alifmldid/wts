package server

import (
	"wts/user"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	userRepo := user.NewUserRepository()
	userUsecase := user.NewUserUsecase(userRepo)
	userController := user.NewUserController(userUsecase)

	registerUserRoute(r, userController)
}