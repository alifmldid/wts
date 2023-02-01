package user

import "context"

type UserUsecase interface{
	UserRegister(c context.Context, payload RegisterPayload) (user User, err error)
	UserLogin()
}

type userUsecase struct{
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase{
	return &userUsecase{userRepository}
}

func (uc *userUsecase) UserRegister(c context.Context, payload RegisterPayload) (user User, err error){
	user, err = uc.userRepository.UserRegister(c, payload)

	return user, err
}

func (uc *userUsecase) UserLogin() (){
	return
}
