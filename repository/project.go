package repository

import (
    "gorm.io/gorm"
	"main.go/model"
)

type ProjectRepository interface {
    Create(project *model.Project) (*model.Project, error)
    GetByID(id int) (*model.Project, error)
    Update(id int, project *model.Project) error
    Delete(id int) error
}

type projectRepository struct{
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *projectRepository {
	return &projectRepository{db}
}

func (p *projectRepository) Create(project *model.Project) (*model.Project, error) {
	if err := p.db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (p *projectRepository) GetByID(id int) (*model.Project, error) {
	var project model.Project
    if err := p.db.First(&project, id).Error; err != nil {
        return nil, err
    }
    return &project, nil
}

func (p *projectRepository) GetAll() ([]*model.Project, error) {
    var projects []*model.Project
    if err := p.db.Find(&projects).Error; err != nil {
        return nil, err
    }
    return projects, nil
}


func (p *projectRepository) Update(id int, project *model.Project) error {
	if err := p.db.Model(&model.Project{}).Where("id = ?", id).Updates(project).Error; err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) Delete(id int) error {
	if err := p.db.Where("id = ?", id).Delete(&model.Project{}).Error; err != nil {
		return err
	}
	return nil
}

