package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

var baseURL = "http://localhost:8001/user"

type UserRepository interface{
	UserRegister(c context.Context, payload RegisterPayload) (user User, err error)
	UserLogin()
}

type userRepository struct{
}

func NewUserRepository() UserRepository{
	return &userRepository{}
}

func (repo *userRepository) UserRegister(c context.Context, payload RegisterPayload) (user User, err error) {
	var responseData Response
	client := &http.Client{}
	userByte, err := json.Marshal(payload)
	if err != nil {
		return User{}, err
	}

	request, err := http.NewRequest("POST", baseURL+"/register", bytes.NewBuffer(userByte))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return User{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return User{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return User{}, err
	}

	user = responseData.Data

	return user, nil
}

func (repo *userRepository) UserLogin() {
	
}