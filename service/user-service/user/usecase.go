package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface{
	Register(c context.Context, payload RegisterPayload) (id string, err error)
	Login(c context.Context, payload LoginPayload) (signedToken string, err error)
	GetUser(c context.Context, id string) (user UserData, err error)
}

type userUsecase struct{
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase{
	return &userUsecase{userRepository}
}

func (uc *userUsecase) Register(c context.Context, payload RegisterPayload) (id string, err error){
	id, err = uc.userRepository.Save(c, payload)

	return id, err
}

func (uc *userUsecase) Login(c context.Context, payload LoginPayload) (signedToken string, err error){
	userData, err := uc.userRepository.GetUserByEmail(c, payload.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(payload.Password))

	if err != nil {
		return "", err
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Id: userData.Id,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	sigedToken, err := token.SignedString([]byte("SECRET_NUMBER"))

	if err != nil {
		return "", err
	}

	return sigedToken, nil
}

func (uc *userUsecase) GetUser(c context.Context, id string) (user UserData, err error){
	user, err = uc.userRepository.GetUserById(c, id)

	return user, err
}