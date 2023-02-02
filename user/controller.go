package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface{
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
}

type userController struct{
	userUsecase UserUsecase
}

func NewUserController(userUsecase UserUsecase) UserController{
	return &userController{userUsecase}
}

func (controller *userController) UserRegister(c *gin.Context){
	var payload RegisterPayload
	c.ShouldBindJSON(&payload)

	user, err := controller.userUsecase.Register(c, payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": user,
	})
}

func (controller *userController) UserLogin(c *gin.Context){
	var payload LoginPayload
	c.ShouldBindJSON(&payload)

	user, err := controller.userUsecase.Login(c, payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": user,
	})
}