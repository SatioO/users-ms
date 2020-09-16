package domain

import "time"

// Book ...
type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Authors   []User    `json:"authors"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
}

// BookUsecase represent the book's usecases
type BookUsecase interface {
	FindBooks() ([]Book, error)
	FindBook(id string) (Book, error)
}

// BookRepository represent the book's repository contract
type BookRepository interface {
	FindAll() ([]Book, error)
	FindOne(id string) (Book, error)
}
