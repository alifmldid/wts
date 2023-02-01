package main

import (
	"user-service/server"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main(){
	r := gin.Default()

	server.RegisterAPIService(r)

	r.Run(":8001")
}