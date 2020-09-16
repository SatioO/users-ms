package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	_userDelivery "github.com/satioO/users/app/users/delivery/http"
	_userRepository "github.com/satioO/users/app/users/repository/mongo"
	_userUsecase "github.com/satioO/users/app/users/usecase"

	_bookDelivery "github.com/satioO/users/app/books/delivery/http"
	_bookRepository "github.com/satioO/users/app/books/repository/mongo"
	_bookUsecase "github.com/satioO/users/app/books/usecase"

	_authDelivery "github.com/satioO/users/app/auth/delivery/http"
	_authRepository "github.com/satioO/users/app/auth/repository"
	_authUsecase "github.com/satioO/users/app/auth/usecase"

	"github.com/satioO/users/app/config"
)

// Server defines the routing configuration
type Server struct {
	Router *mux.Router
	DB     *mongo.Database
}

// Initialize the storage and router
func (a *Server) Initialize() {
	a.DB = config.ConnectDB("users_ms")
	a.Router = mux.NewRouter()

	// Users
	usersRepo := _userRepository.NewUsersRepository(a.DB)
	usersUsecase := _userUsecase.NewUsersUsecase(usersRepo)
	_userDelivery.MakeUserHandlers(a.Router, usersUsecase)

	// Books
	booksRepo := _bookRepository.NewBookRepository(a.DB)
	booksUsecase := _bookUsecase.NewBookUsecase(booksRepo)
	_bookDelivery.MakeBookHandlers(a.Router, booksUsecase)

	// Auth
	authRepo := _authRepository.NewAuthRepository(a.DB)
	authUsecase := _authUsecase.NewAuthUsecase(authRepo)
	_authDelivery.MakeAuthHandlers(a.Router, authUsecase)
}

// Run the application
func (a *Server) Run(port string) {
	fmt.Println("User Management Service is Up Now")
	log.Fatal(http.ListenAndServe(port, a.Router))
}
