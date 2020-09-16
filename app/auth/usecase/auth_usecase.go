package usecase

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/satioO/users/app/domain"
)

type authUsecase struct {
	repo domain.AuthRepository
}

// NewAuthUsecase ...
func NewAuthUsecase(repo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		repo,
	}
}

func (a *authUsecase) LoginUser(auth domain.Auth) (string, error) {
	user, err := a.repo.FindByUsernameAndPassword(auth)

	if err != nil {
		return "", errors.New("Please provide valid login details")
	}

	return a.CreateToken(user)
}

func (a *authUsecase) CreateToken(user domain.User) (string, error) {
	
	os.Setenv("ACCESS_SECRET", "brain_machine")
	claim := jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"expire":  time.Now().Add(time.Minute * 10).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}
