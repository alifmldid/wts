package user

import "time"

type RegisterPayload struct{
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type User struct{
	Id string
	Email string
	Whatsapp string
	Password string
	RegisteredOn time.Time
	UpdatedOn time.Time
}

type Response struct{
	Id string
}