package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/satioO/users/app/middlewares"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"github.com/satioO/users/app/domain"
)

// UsersHandler represent the httphandler for users
type UsersHandler struct {
	usecase domain.UsersUsecase
}

// MakeUserHandlers ...
func MakeUserHandlers(r *mux.Router, usecase domain.UsersUsecase) {
	handler := &UsersHandler{
		usecase,
	}

	r.HandleFunc("/users", middlewares.AuthMiddleware(handler.FindUsers)).Methods("GET")
	r.HandleFunc("/users", middlewares.AuthMiddleware(handler.CreateUser)).Methods("POST")
	r.HandleFunc("/users/{userID}", middlewares.AuthMiddleware(handler.FindUser)).Methods("GET")
	r.HandleFunc("/users/{userID}", middlewares.AuthMiddleware(handler.UpdateUser)).Methods("PUT")
}

// FindUsers will fetch the article based on given params
func (u *UsersHandler) FindUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.usecase.FindUsers()
	if err != nil {
		log.Fatal(err)
	}

	response, _ := json.Marshal(map[string][]domain.User{"data": users})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// FindUser will fetch the user based on given params
func (u *UsersHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	user, err := u.usecase.FindUser(userID)

	if err != nil {
		log.Fatal(err)
	}

	response, _ := json.Marshal(map[string]domain.User{"data": user})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateUser will save the user to db
func (u *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_ = json.NewDecoder(r.Body).Decode(&user)

	err := u.usecase.CreateUser(user)

	if err != nil {
		w.Write([]byte("Failed to write user"))
	} else {
		w.Write([]byte("User saved successfully"))
	}
}

// UpdateUser will update the existing user to db
func (u *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	userID, _ := mux.Vars(r)["userID"]

	_ = json.NewDecoder(r.Body).Decode(&user)

	err := u.usecase.UpdateUser(userID, user)

	if err != nil {
		w.Write([]byte("Failed to update user"))
	} else {
		w.Write([]byte("Successfully updated user"))
	}
}
