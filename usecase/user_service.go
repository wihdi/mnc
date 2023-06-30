package usecase

import (
	"github.com/wihdi/mnc/domain"
   "github.com/wihdi/mnc/repository"
)

type UserUsecase interface {

	FindByUsername(username string) (*domain.User, error)

}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}



func (u *userUsecase) FindByUsername(username string) (*domain.User, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
