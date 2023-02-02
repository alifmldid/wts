package server

import (
	"order-service/config"
	"order-service/order"

	"github.com/gin-gonic/gin"
)

func ServerAPIRegister(r *gin.Engine){
	db := config.GetDBConnection()

	orderRepo := order.NewOrderRepository(db)
	orderUsecase := order.NewOrderUsecase(orderRepo)
	orderController := order.NewOrderController(orderUsecase)

	registerOrderRoute(r, orderController)
}