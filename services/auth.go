package services

import "github.com/steve-kaufman/react-blog-api/storage"

// AuthService is the service for authentication
type AuthService struct {
	userRepo storage.UserRepository
}

// Login takes an email and password and returns a token or an error
func (service *AuthService) Login(email string, password string) (token string, err error) {
	user, err := service.userRepo.GetByEmail(email)

	if err != nil {
		return "", ErrUserNotFound
	}

	if user.Password != password {
		return "", ErrBadPassword
	}

	return "token", nil
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo storage.UserRepository) *AuthService {
	service := new(AuthService)

	service.userRepo = userRepo

	return service
}
