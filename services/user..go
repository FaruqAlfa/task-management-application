package services

import (
    "main.go/model"
    "main.go/repository"
)

type UserService interface {
    Create(user *model.User) (*model.User, error)
    GetByID(id int) (*model.User, error)
    Update(id int, user *model.User) error
    Delete(id int) error
}

type userServices struct {
    userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
    return &userServices{userRepository}
}

func (u *userServices) Create(user *model.User) (*model.User, error) {
   user, err := u.userRepository.Create(user)
   if err != nil {
       return nil, err
   }
   return user, nil
}

func (u *userServices) GetByID(id int) (*model.User, error) {
    user, err := u.userRepository.GetByID(id)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (u *userServices) Update(id int, user *model.User) error {
    err := u.userRepository.Update(id, user)
    if err != nil {
        return err
    }
    return nil
}

func (u *userServices) Delete(id int) error {
    err := u.userRepository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}



