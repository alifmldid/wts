package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface{
	InsertOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	StatusUpdate(c *gin.Context)
}

type orderController struct{
	orderUsecase OrderUsecase
}

func NewOrderController(orderUsecase OrderUsecase) OrderController{
	return &orderController{orderUsecase}
}

func (controller *orderController) InsertOrder(c *gin.Context){
	var payload OrderPayload
	var order NewOrder
	c.ShouldBindJSON(&payload)
	
	id, err := controller.orderUsecase.Insert(c, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	order.Id = id
	c.JSON(200, gin.H{
		"message": "success",
		"data": order,		
	})
}

func (controller *orderController) GetOrder(c *gin.Context){
	id := c.Param("id")

	order, err := controller.orderUsecase.GetOrder(c, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": order,
	})
}

func (controller *orderController) StatusUpdate(c *gin.Context){
	id := c.Param("id")
	status := c.Param("status")

	err := controller.orderUsecase.StatusUpdate(c, id, status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}