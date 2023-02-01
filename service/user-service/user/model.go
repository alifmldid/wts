package user

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type RegisterPayload struct{
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type LoginPayload struct{
	Email string `json:"email"`
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

type NewUserResponse struct{
	Id string
}

type LoginResponse struct{
	Token string
}


type MyClaims struct {
	jwt.StandardClaims
	Id string `json:"id"`
}