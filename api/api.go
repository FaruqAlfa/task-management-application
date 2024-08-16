// api/api.go
package api

import (
	"fmt"
	"net/http"
	"main.go/services"
)

type API struct {
	userService    services.UserService
	taskService    services.TaskService
	projectService services.ProjectService
	adminService   services.AdminService
	mux            *http.ServeMux
}

func NewAPI(userService services.UserService, taskService services.TaskService, projectService services.ProjectService, adminService services.AdminService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		taskService,
		projectService,
		adminService,
		mux,
	}
	
	// User routes
	mux.HandleFunc("/users", api.handleRequest(api.CreateUser, http.MethodPost))
	mux.HandleFunc("/users/get", api.handleRequest(api.GetUser, http.MethodGet))
	mux.HandleFunc("/users/update", api.handleRequest(api.UpdateUser, http.MethodPut))
	mux.HandleFunc("/users/delete", api.handleRequest(api.DeleteUser, http.MethodDelete))

	// Task routes
	mux.HandleFunc("/tasks", api.handleRequest(api.GetTask, http.MethodGet))        
	mux.HandleFunc("/tasks/add", api.handleRequest(api.CreateTask, http.MethodPost))  
	mux.HandleFunc("/tasks/update", api.handleRequest(api.UpdateTask, http.MethodPut)) 
	mux.HandleFunc("/tasks/delete", api.handleRequest(api.DeleteTask, http.MethodDelete)) 

	// Project routes
	mux.HandleFunc("/project", api.handleRequest(api.GetProject, http.MethodGet))       
	mux.HandleFunc("/project/add", api.handleRequest(api.CreateProject, http.MethodPost)) 
	mux.HandleFunc("/project/update", api.handleRequest(api.UpdateProject, http.MethodPut)) 
	mux.HandleFunc("/project/delete", api.handleRequest(api.DeleteProject, http.MethodDelete)) 

	// Admin routes
	mux.HandleFunc("/admin/register", api.handleRequest(api.CreateAdmin, http.MethodPost))
	mux.HandleFunc("/admin/login", api.handleRequest(api.LoginAdmin, http.MethodPost))

	return api
}


func (api *API) handleRequest(fn http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fn(w, r)
	}
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
