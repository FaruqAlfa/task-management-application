package main

import (
	"main.go/api"
	"main.go/db"
	"main.go/model"
	repo "main.go/repository"
	"main.go/services"

	"log"
)

func main() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "taskmanagement",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Task{}, &model.Project{}, &model.Admin{})
	err = conn.AutoMigrate(&model.Admin{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	userRepo := repo.NewUserRepo(conn)
	taskRepo := repo.NewTaskRepo(conn)
	projectRepo := repo.NewProjectRepo(conn)
	adminRepo := repo.NewAdminRepository(conn)

	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo, userRepo, projectRepo)
	projectService := services.NewProjectService(projectRepo)
	adminService := services.NewAdminService(adminRepo)

	mainAPI := api.NewAPI(userService, taskService, projectService, adminService)
	mainAPI.Start()
}
