package main

import (
	"wts/server"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main(){
	r := gin.Default()

	server.RegisterAPIService(r)

	r.Run(":8000")
}