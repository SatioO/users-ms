package domain

// Auth ...
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AuthUsecase ...
type AuthUsecase interface {
	LoginUser(auth Auth) (string, error)
	CreateToken(user User) (string, error)
}

// AuthRepository ...
type AuthRepository interface {
	FindByUsernameAndPassword(auth Auth) (User, error)
}
