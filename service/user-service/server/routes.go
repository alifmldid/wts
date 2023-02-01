package server

import (
	"user-service/user"

	"github.com/gin-gonic/gin"
)

func registerUserRoute(r *gin.Engine, userController user.UserController){
	user := r.Group("/user")
	user.POST("/register", userController.UserRegister)
	user.POST("/login", userController.UserLogin)
}