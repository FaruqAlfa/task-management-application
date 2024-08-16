package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"main.go/model"
	
)

func (api *API) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    createdUser, err := api.userService.Create(&user) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser) 
}

func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }
    
    idInt, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    if err := api.userService.Update(idInt, &user); err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}


func (api *API) GetUser(w http.ResponseWriter, r *http.Request) {
   id := r.URL.Query().Get("id")
   idInt, err := strconv.Atoi(id)
   if err != nil {
       http.Error(w, err.Error(), http.StatusBadRequest)
       return
   }
   user, err := api.userService.GetByID(idInt)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }
   w.WriteHeader(http.StatusOK)
   json.NewEncoder(w).Encode(user)
}


func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
   id := r.URL.Query().Get("id")
   idInt, err := strconv.Atoi(id)
   if err != nil {
       http.Error(w, err.Error(), http.StatusBadRequest)
       return
   }
   if err := api.userService.Delete(idInt); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }
   w.WriteHeader(http.StatusOK)
   json.NewEncoder(w).Encode(model.SuccessResponse{Message: "user berhasil dihapus"})
}