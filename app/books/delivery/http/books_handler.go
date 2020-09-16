package http

import (
	"encoding/json"
	"net/http"

	"github.com/satioO/users/app/middlewares"

	"github.com/gorilla/mux"
	"github.com/satioO/users/app/domain"
)

type bookHandler struct {
	usecase domain.BookUsecase
}

// MakeBookHandlers ...
func MakeBookHandlers(r *mux.Router, usecase domain.BookUsecase) {
	handler := &bookHandler{
		usecase,
	}

	r.HandleFunc("/books", middlewares.AuthMiddleware(handler.FindBooks)).Methods("GET")
	r.HandleFunc("/book/{bookID}", middlewares.AuthMiddleware(handler.FindBook)).Methods("GET")
}

func (b *bookHandler) FindBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := b.usecase.FindBooks()

	response, _ := json.Marshal(map[string][]domain.Book{"data": books})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func (b *bookHandler) FindBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["bookID"]

	book, _ := b.usecase.FindBook(bookID)

	response, _ := json.Marshal(map[string]domain.Book{"data": book})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}
