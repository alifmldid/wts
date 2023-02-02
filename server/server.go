package server

import (
	"wts/order"
	"wts/ticket"
	"wts/user"

	"github.com/gin-gonic/gin"
)

func RegisterAPIService(r *gin.Engine){
	userRepo := user.NewUserRepository()
	userUsecase := user.NewUserUsecase(userRepo)
	userController := user.NewUserController(userUsecase)

	registerUserRoute(r, userController)

	ticketRepo := ticket.NewTicketRepository()
	ticketUsecase := ticket.NewTicketUsecase(ticketRepo)
	ticketController := ticket.NewTicketController(ticketUsecase)

	registerTicketRoute(r, ticketController)

	orderRepo := order.NewOrderRepository()
	orderUsecase := order.NewOrderUsecase(orderRepo)
	orderController := order.NewOrderController(orderUsecase)

	registerOrderRoute(r, orderController)
}