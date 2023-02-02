package order

import "time"

type Order struct{
	Id string
	UserID string
	TicketID string
	Qty int
	Status string
	RegisteredOn time.Time
	UpdatedOn time.Time
}

type NewOrder struct{
	Id string `json:"id"`
}

type OrderData struct{
	Id string `json:"id"`
	User User `json:"buyer"`
	Ticket Ticket `json:"ticket"`
	Qty int `json:"qty"`
	Status string `json:"status"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type User struct{
	ID string `json:"id"`
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`	
}

type Ticket struct{
	ID string `json:"id"`
	User Owner `json:"owner"`	
	Event string `json:"event"`
	Qty int `json:"qty"`
	Status string `json:"status"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type Owner struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type OrderPayload struct{
	TicketID string `json:"ticket_id"`
	Qty int `json:"qty"`
}

type TicketDataResponse struct{
	Message string `json:"message"`
	Data Ticket `json:"data"`
}

type UserDataResponse struct{
	Message string `json:"message"`
	Data User `json:"data"`
}