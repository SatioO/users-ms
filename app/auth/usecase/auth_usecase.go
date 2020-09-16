package usecase

import (
	"os"
	"time"

	"github.com/twinj/uuid"

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

func (a *authUsecase) LoginUser(auth domain.Auth) (*domain.TokenDetails, error) {
	user, err := a.repo.FindByUsernameAndPassword(auth)

	if err != nil {
		return nil, errors.New("Please provide valid login details")
	}

	return a.CreateToken(user)
}

func (a *authUsecase) CreateToken(user domain.User) (*domain.TokenDetails, error) {
	td := &domain.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 60).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Minute * 300).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error

	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "brain_machine")
	atClaims := jwt.MapClaims{
		"user_id":     user.ID,
		"username":    user.Name,
		"exp":         td.AtExpires,
		"access_uuid": td.AccessUUID,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	// Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "machine")
	rtClaims := jwt.MapClaims{
		"user_id":      user.ID,
		"username":     user.Name,
		"exp":          td.RtExpires,
		"refresh_uuid": td.RefreshUUID,
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {
		return nil, err
	}

	return td, nil
}
