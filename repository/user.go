package repository

import (
	"gorm.io/gorm"
	"main.go/model"
)

type UserRepository interface {
    Create(user *model.User) (*model.User, error)
    GetByID(id int) (*model.User, error)
    Update(id int, user *model.User) error
    Delete(id int) error
	GetByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// repository/userRepo.go
func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (u *userRepository) Update(id int, user *model.User) error {
	if err := u.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id int) error {
	if err := u.db.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}


func (u *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
