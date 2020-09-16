package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satioO/users/app/domain"
)

type authHandler struct {
	usecase domain.AuthUsecase
}

// MakeAuthHandlers ...
func MakeAuthHandlers(r *mux.Router, usecase domain.AuthUsecase) {
	handler := &authHandler{
		usecase,
	}

	r.HandleFunc("/auth/login", handler.LoginUser).Methods("POST")
}

func (a *authHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var auth domain.Auth

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid JSON provided"))
		return
	}

	tokenDetails, err := a.usecase.LoginUser(auth)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(map[string]domain.TokenDetails{"data": *tokenDetails})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
