package order

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type OrderRepository interface{
	Save(c context.Context, payload OrderPayload) (id string, err error)
	GetById(c context.Context, id string) (orderData OrderData, err error)
	Update(c context.Context, id string, status string) (err error)
	Export(c context.Context, email string) (err error)
}

type orderRepository struct{
	Conn *gorm.DB
}

func NewOrderRepository(Conn *gorm.DB) OrderRepository{
	return &orderRepository{Conn}
}

func (repo *orderRepository) Save(c context.Context, payload OrderPayload) (id string, err error){
	var order Order

	randId, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
	order.Id = "order-"+randId
	userInfo := c.Value("userInfo").(jwt.MapClaims)
	order.UserID = userInfo["id"].(string)
	order.TicketID = payload.TicketID
	order.Qty = payload.Qty
	order.Status = "not paid"
	order.RegisteredOn = time.Now()
	order.UpdatedOn = time.Now()

	err = repo.Conn.Save(&order).Error
	if err != nil {
		return "", err
	}

	return order.Id, nil
}

func (repo *orderRepository) GetById(c context.Context, id string) (orderData OrderData, err error){
	var order Order
	err = repo.Conn.Where("id = ?", id).First(&order).Error
	if err != nil {
		return OrderData{}, err
	}

	client := &http.Client{}

	requestTicket, err := http.NewRequest("GET", "http://localhost:8002/ticket/"+order.TicketID, nil)
	if err != nil {
		return OrderData{}, err
	}

	responseTicket, err := client.Do(requestTicket)	
	if err != nil {
		return OrderData{}, err
	}
	defer responseTicket.Body.Close()

	var responseTicketData TicketDataResponse
	err = json.NewDecoder(responseTicket.Body).Decode(&responseTicketData)
	if err != nil {
		return OrderData{}, err
	}

	var responseUserData UserDataResponse

	requestUser, err := http.NewRequest("GET", "http://localhost:8001/user/"+order.UserID, nil)
	if err != nil {
		return OrderData{}, err
	}

	responseUser, err := client.Do(requestUser)
	if err != nil {
		return OrderData{}, err
	}
	defer responseUser.Body.Close()

	err = json.NewDecoder(responseUser.Body).Decode(&responseUserData)
	if err != nil {
		return OrderData{}, err
	}

	orderData.Id = order.Id
	orderData.Ticket = responseTicketData.Data
	orderData.User = responseUserData.Data
	orderData.Qty = order.Qty
	orderData.Status = order.Status
	orderData.RegisteredOn = order.RegisteredOn
	orderData.UpdatedOn = order.UpdatedOn

	return orderData, nil
}

func (repo *orderRepository) Update(c context.Context, id string, status string) (err error){
	err = repo.Conn.Where("id = ?", id).Updates(Order{Status: status}).Error

	return err
}

func (repo *orderRepository) Export(c context.Context, email string) (err error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"paid-notif",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body := email

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}