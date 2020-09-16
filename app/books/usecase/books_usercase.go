package usecase

import (
	"github.com/satioO/users/app/domain"
)

type bookUsecase struct {
	repo domain.BookRepository
}

// NewBookUsecase ...
func NewBookUsecase(repo domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		repo,
	}
}

func (b *bookUsecase) FindBooks() ([]domain.Book, error) {
	return b.repo.FindAll()
}

func (b *bookUsecase) FindBook(id string) (domain.Book, error) {
	return b.repo.FindOne(id)
}
