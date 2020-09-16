package usecase

import (
	"github.com/satioO/users/app/domain"
)

type usersUsecase struct {
	repo domain.UserRepository
}

// NewUsersUsecase will create new an usersUsecase object representation of domain.UsersUsecase interface
func NewUsersUsecase(repo domain.UserRepository) domain.UsersUsecase {
	return &usersUsecase{
		repo,
	}
}

func (u *usersUsecase) FindUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *usersUsecase) FindUser(id string) (domain.User, error) {
	return u.repo.FindOne(id)
}

func (u *usersUsecase) CreateUser(user domain.User) error {
	return u.repo.Save(user)
}

func (u *usersUsecase) UpdateUser(userID string, user domain.User) error {
	return u.repo.UpdateOne(userID, user)
}
