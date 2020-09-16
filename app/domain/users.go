package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updatedAt"`
}

// UsersUsecase represent the user's usecases
type UsersUsecase interface {
	FindUsers() ([]User, error)
	FindUser(id string) (User, error)
	CreateUser(user User) error
	UpdateUser(userID string, user User) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	FindAll() (res []User, err error)
	FindOne(id string) (User, error)
	Save(user User) error
	UpdateOne(userID string, user User) error
}
