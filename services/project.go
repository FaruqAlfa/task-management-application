package services

import (
    "main.go/model"
    "main.go/repository"
)

type ProjectService interface {
    Create(project *model.Project) (*model.Project, error)
    GetByID(id int) (*model.Project, error)
    Update(id int, project *model.Project) error
    Delete(id int) error
}

type projectService struct {
    projectRepository repository.ProjectRepository
}

func NewProjectService(projectRepository repository.ProjectRepository) *projectService {
    return &projectService{projectRepository}
}

func (p *projectService) Create(project *model.Project) (*model.Project, error) {
    project, err := p.projectRepository.Create(project)
    if err != nil {
        return nil, err
    }
    return project, nil
}


func (p *projectService) GetByID(id int) (*model.Project, error) {
  project, err := p.projectRepository.GetByID(id)
  if err != nil {
    return nil, err
  }

  return project, nil
}

func (p *projectService) Update(id int, project *model.Project) error {
  err := p.projectRepository.Update(id, project)
  if err != nil {
    return err
  }

  return nil
}

func (p *projectService) Delete(id int) error {
  err := p.projectRepository.Delete(id)
  if err != nil {
    return err
  }

  return nil
}
