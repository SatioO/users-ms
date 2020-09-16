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

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     string
	Username   string
}

// AuthUsecase ...
type AuthUsecase interface {
	LoginUser(auth Auth) (*TokenDetails, error)
	CreateToken(user User) (*TokenDetails, error)
}

// AuthRepository ...
type AuthRepository interface {
	FindByUsernameAndPassword(auth Auth) (User, error)
}
