package ticket

import "time"

type TicketPayload struct{
	Event string `json:"event"`
	Qty int `json:"qty"`
}

type Ticket struct{
	Id string
	UserID string
	Event string
	Qty int
	Status string
	IsDeleted bool
	RegisteredOn time.Time
	UpdatedOn time.Time
}

type NewTicketResponse struct{
	Id string `json:"id"`
}

type TicketData struct{
	Id string `json:"id"`
	User UserData `json:"owner"`
	Event string `json:"event"`
	Qty int `json:"qty"`
	Status string `json:"status"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type UserDataResponse struct{
	Message string `json:"message"`
	Data UserData `json:"data"`
}

type UserData struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	RegisteredOn time.Time `json:"registered_on"`
	UpdatedOn time.Time `json:"updated_on"`		
}
