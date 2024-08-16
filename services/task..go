package services

import (
	"fmt"
	"main.go/model"
	"main.go/repository"
)

type TaskService interface {
	Create(task *model.Task) (*model.Task, error)
	GetByID(id int) (*model.Task, error)
	Update(id int, task *model.Task) error
	Delete(id int) error
	ValidateUserID(userID int) error
   ValidateProjectID(projectID uint) error
}

type taskService struct {
	taskRepository repository.TaskRepository
	userRepository repository.UserRepository 
   projectRepository repository.ProjectRepository
}

func NewTaskService(taskRepository repository.TaskRepository, userRepository repository.UserRepository, projectRepository repository.ProjectRepository) TaskService {
   return &taskService{
       taskRepository:    taskRepository,
       userRepository:    userRepository,
       projectRepository: projectRepository, 
   }
}

func (t *taskService) Create(task *model.Task) (*model.Task, error) {
	task, err := t.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) GetByID(id int) (*model.Task, error) {
	task, err := t.taskRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) Update(id int, task *model.Task) error {
	err := t.taskRepository.Update(id, task)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) Delete(id int) error {
	err := t.taskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) ValidateUserID(userID int) error {
	_, err := t.userRepository.GetByID(userID)
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}
	return nil
}

func (t *taskService) ValidateProjectID(projectID uint) error {
   project, err := t.projectRepository.GetByID(int(projectID)) 
   if err != nil || project == nil {
       return fmt.Errorf("project with ID %d does not exist", projectID)
   }
   return nil
}

