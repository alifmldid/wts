package ticket

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

var baseURL = "http://localhost:8002/ticket"

type TicketRepository interface{
	Save(c context.Context, payload TicketPayload) (ticket NewTicket, err error)
	GetTicketById(c context.Context, id string) (ticket Ticket, err error)
	DataUpdate(c context.Context, id string, payload TicketPayload) (err error)
	StatusUpdate(c context.Context, id string, status string) (err error)
}

type ticketRepository struct{

}

func NewTicketRepository() TicketRepository{
	return &ticketRepository{}
}

func (repo *ticketRepository) Save(c context.Context, payload TicketPayload) (ticket NewTicket, err error){
	var responseData InsertResponse
	client := &http.Client{}

	ticketByte, err := json.Marshal(payload)
	if err != nil {
		return NewTicket{}, err
	}

	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(ticketByte))
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		return NewTicket{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return NewTicket{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return NewTicket{}, err
	}

	ticket = responseData.Data

	return ticket, nil
}

func (repo *ticketRepository) GetTicketById(c context.Context, id string) (ticket Ticket, err error){
	var responseData GetResponse
	client := &http.Client{}

	request, err := http.NewRequest("GET", baseURL+"/"+id, nil)
	if err != nil {
		return Ticket{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return Ticket{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return Ticket{}, err
	}

	ticket = responseData.Data

	return
}

func (repo *ticketRepository) DataUpdate(c context.Context, id string, payload TicketPayload) (err error){
	client := &http.Client{}

	ticketByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PUT", baseURL+"/"+id, bytes.NewBuffer(ticketByte))
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (repo *ticketRepository) StatusUpdate(c context.Context, id string, status string) (err error){
	client := &http.Client{}

	request, err := http.NewRequest("PUT", baseURL+"/"+id+"/"+status, nil)
	token := c.Value("token").(string)
	request.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}