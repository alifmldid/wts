package user

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface{
	UserRegister(c context.Context, payload RegisterPayload) (id string, err error)
	UserLogin(c context.Context, user User)
}

type userRepository struct{
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) UserRepository{
	return &userRepository{Conn}
}

func (repo *userRepository) UserRegister(c context.Context, payload RegisterPayload) (id string, err error){
	var user User

	randId, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
	if err != nil {
		return "", err
	}

	user.Id = "user-"+randId
	user.Email = payload.Email
	user.Whatsapp = payload.Whatsapp

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	user.Password = string(hashedPassword)

	user.RegisteredOn = time.Now()
	user.UpdatedOn = time.Now()

	createdUser := repo.Conn.Create(&user).Error
	if createdUser != nil {
		return "", err
	}
	return user.Id, nil
}

func (repo *userRepository) UserLogin(c context.Context, user User){

}