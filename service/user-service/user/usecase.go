package user

import "context"

type UserUsecase interface{
	UserRegister(c context.Context, payload RegisterPayload) (id string, err error)
	UserLogin()
}

type userUsecase struct{
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase{
	return &userUsecase{userRepository}
}

func (uc *userUsecase) UserRegister(c context.Context, payload RegisterPayload) (id string, err error){
	id, err = uc.userRepository.UserRegister(c, payload)

	return id, err
}

func (uc *userUsecase) UserLogin(){

}