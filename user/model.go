package user

type RegisterPayload struct{
	Email string `json:"email"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type LoginPayload struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type NewUserResponse struct{
	Message string `json:"message"`
	Data NewUser `json:"data"`
}

type NewUser struct{
	Id string `json:"id"`
}

type UserResponse struct{
	Message string `json:"message"`
	Data User `json:"data"`
}

type User struct{
	Token string `json:"token"`
}