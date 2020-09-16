package repository

import (
	"context"

	"github.com/satioO/users/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	db *mongo.Database
}

// NewAuthRepository ...
func NewAuthRepository(db *mongo.Database) domain.AuthRepository {
	return &authRepository{
		db,
	}
}

func (a *authRepository) FindByUsernameAndPassword(auth domain.Auth) (domain.User, error) {
	var user domain.User

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"name": auth.Username},
			bson.M{"password": auth.Password},
		},
	}

	err := a.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)

	return user, err
}
