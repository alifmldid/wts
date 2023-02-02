package ticket

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketController interface{
	TicketInsert(c *gin.Context)
	GetTicket(c *gin.Context)
	DataUpdate(c *gin.Context)
	StatusUpdate(c *gin.Context)
}

type ticketController struct{
	ticketUsecase TicketUsecase
}

func NewTicketController(ticketUsecase TicketUsecase) TicketController{
	return &ticketController{ticketUsecase}
}

func (controller *ticketController) TicketInsert(c *gin.Context){
	var payload TicketPayload
	c.ShouldBindJSON(&payload)

	ticket, err := controller.ticketUsecase.Insert(c, payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": ticket,
	})
}

func (controller *ticketController) GetTicket(c *gin.Context){
	id := c.Param("id")

	ticket, err := controller.ticketUsecase.GetTicket(c, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": ticket,
	})
}

func (controller *ticketController) DataUpdate(c *gin.Context) {
	id := c.Param("id")
	var payload TicketPayload
	c.ShouldBindJSON(&payload)

	err := controller.ticketUsecase.DataUpdate(c, id, payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (controller *ticketController) StatusUpdate(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	err := controller.ticketUsecase.StatusUpdate(c, id, status)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}