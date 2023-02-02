package ticket

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	gonanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"
)

type TicketRepository interface{
	Save(c context.Context, payload TicketPayload) (id string, err error)
	GetById(c context.Context, id string) (ticket TicketData, err error)
	DataUpdate(c context.Context, id string, payload TicketPayload) (err error)
	StatusUpdate(c context.Context, id string, status string) (err error)
}

type ticketRepository struct{
	Conn *gorm.DB
}

func NewTicketRepository(Conn *gorm.DB) TicketRepository{
	return &ticketRepository{Conn}
}

func (repo *ticketRepository) Save(c context.Context, payload TicketPayload) (id string, err error){
	var ticket Ticket

	randId, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
	if err != nil {
		return "", err
	}

	userInfo := c.Value("userInfo").(jwt.MapClaims)

	ticket.Id = "ticket-"+randId
	ticket.UserID = userInfo["id"].(string)
	ticket.Event = payload.Event
	ticket.Qty = payload.Qty
	ticket.Status = "active"
	ticket.IsDeleted = false
	ticket.RegisteredOn = time.Now()
	ticket.UpdatedOn = time.Now()

	err = repo.Conn.Save(&ticket).Error

	if err != nil {
		return "", err
	}

	return ticket.Id, err
}

func (repo *ticketRepository) GetById(c context.Context, id string) (ticketData TicketData, err error){
	var ticket Ticket

	err = repo.Conn.Where("id = ?", id).First(&ticket).Error

	if err != nil {
		return TicketData{}, err
	}

	userId := ticket.UserID

	var responseData UserDataResponse
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://localhost:8001/user/"+userId, nil)
	if err != nil {
		return TicketData{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return TicketData{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return TicketData{}, err
	}


	ticketData.Id = ticket.Id
	ticketData.User = UserData(responseData.Data)
	ticketData.Event = ticket.Event
	ticketData.Qty = ticket.Qty
	ticketData.Status = ticket.Status
	ticketData.RegisteredOn = ticket.RegisteredOn
	ticketData.UpdatedOn = ticket.UpdatedOn

	return ticketData, nil
}

func (repo *ticketRepository) DataUpdate(c context.Context, id string, payload TicketPayload) (err error){
	err = repo.Conn.Where("id = ?", id).Updates(Ticket{Event: payload.Event, Qty: payload.Qty, UpdatedOn: time.Now()}).Error

	return err
}

func (repo *ticketRepository) StatusUpdate(c context.Context, id string, status string) (err error){
	err = repo.Conn.Where("id = ?", id).Updates(Ticket{Status: status, UpdatedOn: time.Now()}).Error

	return err
}