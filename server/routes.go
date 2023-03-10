package server

import (
	"fmt"
	"net/http"
	"strings"
	"wts/order"
	"wts/ticket"
	"wts/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func registerUserRoute(r *gin.Engine, userController user.UserController){
	user := r.Group("/user")
	user.POST("/register", userController.UserRegister)
	user.POST("/login", userController.UserLogin)
}

func registerTicketRoute(r *gin.Engine, ticketController ticket.TicketController){
	ticket := r.Group("/")
	ticket.Use(authRequired)
	ticket.POST("/ticket", ticketController.TicketInsert)
	ticket.PUT("/ticket/:id", ticketController.DataUpdate)
	ticket.PUT("/ticket/:id/:status", ticketController.StatusUpdate)
	r.GET("/ticket/:id", ticketController.GetTicket)
}

func registerOrderRoute(r *gin.Engine, orderController order.OrderController){
	order := r.Group("/")
	order.Use(authRequired)
	order.POST("/order", orderController.InsertOrder)
	order.PUT("/order/:id/:status", orderController.StatusUpdate)
	r.GET("/order/:id", orderController.GetOrder)
}

func authRequired(c *gin.Context){
	authorizationHeader := c.Request.Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte("SECRET_NUMBER"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userInfo, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Set("token", token.Raw)
	c.Set("userInfo", userInfo)
	c.Next()
}