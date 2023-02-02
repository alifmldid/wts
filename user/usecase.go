package user

import "context"

type UserUsecase interface{
	Register(c context.Context, payload RegisterPayload) (user NewUser, err error)
	Login(c context.Context, payload LoginPayload) (user User, err error)
}

type userUsecase struct{
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase{
	return &userUsecase{userRepository}
}

func (uc *userUsecase) Register(c context.Context, payload RegisterPayload) (user NewUser, err error){
	user, err = uc.userRepository.Save(c, payload)

	return user, err
}

func (uc *userUsecase) Login(c context.Context, payload LoginPayload) (user User, err error){
	user, err = uc.userRepository.GetUserByEmail(c, payload)

	return user, err
}
