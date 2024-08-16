package api

import (
	"encoding/json"
	"net/http"
	"main.go/model"
)

func (api *API) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.adminService.Create(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

func (api *API) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := api.adminService.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
