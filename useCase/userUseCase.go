package useCase

import (
	"github.com/nattapat27/test-golnag-echo/model"
	"github.com/nattapat27/test-golnag-echo/repository"
)

type userUseCase struct {
	userRepo repository.UserRepositoryInf
}

func NewUserUseCase(userRepo repository.UserRepositoryInf) UserUseCaseInf {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Create(user *model.User) error {
	return u.userRepo.Create(user)
}

func (u userUseCase) FetchOne(id int) (*model.User, error) {
	return u.userRepo.FetchOne(id)
}

func (u *userUseCase) FetchAll() ([]*model.User, error) {
	return u.userRepo.Fetch()
}