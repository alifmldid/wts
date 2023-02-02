package server

import (
	"ticket-service/config"
	"ticket-service/ticket"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine) {
	db := config.GetDBConnection()

	ticketRepo := ticket.NewTicketRepository(db)
	ticketUsecase := ticket.NewTicketUsecase(ticketRepo)
	ticketController := ticket.NewTicketController(ticketUsecase)

	registerTicketRoute(r, ticketController)
}