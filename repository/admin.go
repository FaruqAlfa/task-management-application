// repository/admin_repository.go
package repository

import (
	"main.go/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
	GetByEmail(email string) (*model.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) GetByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	return &admin, err
}
