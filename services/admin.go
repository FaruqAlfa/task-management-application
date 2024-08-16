package services

import (
	"main.go/model"
	"main.go/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"time"
)

type AdminService interface {
	Create(admin *model.Admin) error
	Login(email, password string) (string, error)
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &adminService{adminRepo}
}

func (s *adminService) Create(admin *model.Admin) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)
	return s.adminRepo.Create(admin)
}

func (s *adminService) Login(email, password string) (string, error) {
	admin, err := s.adminRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"admin_id": admin.ID,
		"email":    admin.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-256-bit-secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
