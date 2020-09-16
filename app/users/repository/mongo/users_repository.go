package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/satioO/users/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersRepository struct {
	db *mongo.Database
}

// NewUsersRepository will create an object that represent the article.Repository interface
func NewUsersRepository(db *mongo.Database) domain.UserRepository {
	return &usersRepository{
		db,
	}
}

// FindAll ...
func (u *usersRepository) FindAll() ([]domain.User, error) {
	cursor, err := u.db.Collection("users").Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	users := []domain.User{}
	err = cursor.All(context.TODO(), &users)

	return users, err
}

// FindOne ...
func (u *usersRepository) FindOne(id string) (domain.User, error) {
	user := domain.User{}
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	err = u.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)

	return user, err
}

// Save ..
func (u *usersRepository) Save(user domain.User) error {
	_, err := u.db.Collection("users").InsertOne(context.TODO(), user)
	return err
}

func (u *usersRepository) UpdateOne(userID string, user domain.User) error {
	id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{
		"$set": bson.M{
			"name": user.Name,
		},
	}
	_, err := u.db.Collection("users").UpdateOne(context.TODO(), filter, update)
	return err
}
