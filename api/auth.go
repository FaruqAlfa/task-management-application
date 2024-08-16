package api

import (
	"encoding/json"
	"net/http"
	"main.go/model"
	"main.go/services"
)

type AuthAPI struct {
	authService services.AuthService
}

func NewAuthAPI(authService services.AuthService) *AuthAPI {
	return &AuthAPI{authService: authService}
}

func (api *AuthAPI) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := api.authService.Login(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
