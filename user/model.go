package user

type RegisterPayload struct{
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type Response struct{
	Message string `json:"message"`
	Data User `json:"data"`
}

type User struct{
	Id string `json:"id"`
}