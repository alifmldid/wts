package order

import "time"

type OrderPayload struct{
	TicketID string `json:"ticket_id"`
	Qty int `json:"qty"`
}

type NewOrderResponse struct{
	Message string
	Data NewOrder
}

type NewOrder struct{
	Id string `json:"id"`
}

type OrderResponse struct{
	Message string
	Data Order
}

type Order struct{
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