package usecase

import (
	"context"

	"github.com/satioO/users/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookRepository struct {
	db *mongo.Database
}

// NewBookRepository ...
func NewBookRepository(db *mongo.Database) domain.BookRepository {
	return &bookRepository{
		db,
	}
}

// FindAll ...
func (b *bookRepository) FindAll() ([]domain.Book, error) {
	cursor, err := b.db.Collection("books").Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	books := []domain.Book{}
	err = cursor.All(context.TODO(), &books)

	return books, err
}

// FindOne ...
func (b *bookRepository) FindOne(id string) (domain.Book, error) {
	return domain.Book{}, nil
}
