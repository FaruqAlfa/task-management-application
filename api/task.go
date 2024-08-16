package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"main.go/model"
	
)

func (api *API) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

  // Konversi uint ke int
  userID := int(task.UserID)

  
  if err := api.taskService.ValidateUserID(userID); err != nil {
	  http.Error(w, "Invalid user ID", http.StatusBadRequest)
	  return
  }


    if err := api.taskService.ValidateProjectID(task.ProjectID); err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }

   
    createdTask, err := api.taskService.Create(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdTask)
}




func (api *API) UpdateTask(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := api.taskService.Update(idInt, &task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}



func (api *API) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := api.taskService.GetByID(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}


func (api *API) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := api.taskService.Delete(idInt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "task deleted successfully"})
}