package ticket

import "time"

type TicketPayload struct{
	Event string `json:"event"`
	Qty int `json:"qty"`
}

type InsertResponse struct{
	Message string `json:"message"`
	Data NewTicket `json:"data"`
}

type NewTicket struct{
	Id string `json:"id"`
}

type GetResponse struct{
	Message string `json:"message"`
	Data Ticket `json:"data"`
}

type Ticket struct{
	Id string `json:"id"`
	User User `json:"owner"`
	Event string `json:"event"`
	Qty int `json:"qty"`
	Status string `json:"status"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type User struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}