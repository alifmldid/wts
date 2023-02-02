package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface{
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
	GetUser(c *gin.Context)
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

	id, err := controller.userUsecase.Register(c, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	var response NewUserResponse
	response.Id = id

	c.JSON(200, gin.H{
		"message": "success",
		"data": response,
	})
}

func (controller *userController) UserLogin(c *gin.Context){
	var payload LoginPayload
	c.ShouldBindJSON(&payload)

	signedToken, err := controller.userUsecase.Login(c, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var response LoginResponse
	response.Token = signedToken

	c.JSON(200, gin.H{
		"message": "success",
		"data": response,
	})
}

func (controller *userController) GetUser(c *gin.Context){
	id := c.Param("id")

	user, err := controller.userUsecase.GetUser(c, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": user,
	})
}