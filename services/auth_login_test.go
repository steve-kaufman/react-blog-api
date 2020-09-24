package services_test

import (
	"testing"

	"github.com/steve-kaufman/react-blog-api/models"
	"github.com/steve-kaufman/react-blog-api/services"
	"github.com/steve-kaufman/react-blog-api/storage"
	"github.com/stretchr/testify/assert"
)

func SetupAuthLoginTest() *services.AuthService {
	repo := new(storage.TestUserRepository)

	repo.Users = []models.User{
		{
			Email:    "123@example.com",
			Password: "password1",
		},
		{
			Email:    "456@example.com",
			Password: "password2",
		},
	}

	return services.NewAuthService(repo)
}

func TestLoginFailsWithInvalidUsername(t *testing.T) {
	authService := SetupAuthLoginTest()

	token, err := authService.Login("johndoe@example.com", "supersecret")

	assert.Empty(t, token)
	assert.Same(t, services.ErrUserNotFound, err)
}

func TestLoginFailsWithInvalidPassword(t *testing.T) {
	authService := SetupAuthLoginTest()

	token, err := authService.Login("123@example.com", "wrongpassword")

	assert.Empty(t, token)
	assert.Same(t, services.ErrBadPassword, err)
}

func TestLoginSucceedsWithValidInfo(t *testing.T) {
	authService := SetupAuthLoginTest()

	token, err := authService.Login("456@example.com", "password2")

	assert.NotEmpty(t, token)
	assert.Nil(t, err)
}
