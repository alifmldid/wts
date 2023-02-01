package server

import (
	"wts/user"

	"github.com/gin-gonic/gin"
)

func registerUserRoute(r *gin.Engine, userController user.UserController){
	user := r.Group("/user")
	user.POST("/register", userController.UserRegister)
}