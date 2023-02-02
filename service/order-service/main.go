package main

import (
	"order-service/server"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	server.ServerAPIRegister(r)

	r.Run(":8003")
}