package services

import (
	"github.com/steve-kaufman/react-blog-api/storage"
	"github.com/steve-kaufman/react-blog-api/util"
)

// AuthService is the service for authentication
type AuthService struct {
	Hasher   util.PasswordHasher
	userRepo storage.UserRepository
}

// Login takes an email and password and returns a token or an error
func (service *AuthService) Login(email string, password string) (token string, err error) {
	user, err := service.userRepo.GetByEmail(email)

	if err != nil {
		return "", ErrUserNotFound
	}

	if user.Password != service.Hasher.Hash(password) {
		return "", ErrBadPassword
	}

	return "token", nil
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo storage.UserRepository) *AuthService {
	service := new(AuthService)

	service.userRepo = userRepo
	service.Hasher = new(util.DefaultHasher)

	return service
}
