package services

import (
	"errors"
	"main.go/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepository: userRepo}
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
	})

	tokenString, err := token.SignedString([]byte("your-256-bit-secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
