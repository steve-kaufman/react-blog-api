package services_test

import (
	"testing"

	"github.com/steve-kaufman/react-blog-api/models"
	"github.com/steve-kaufman/react-blog-api/services"
	"github.com/steve-kaufman/react-blog-api/storage"
	"github.com/steve-kaufman/react-blog-api/util"
	"github.com/stretchr/testify/assert"
)

func SetupAuthLoginTest(usersInDB []models.User) *services.AuthService {
	repo := new(storage.MockUserRepository)
	repo.Users = usersInDB

	authService := services.NewAuthService(repo)
	authService.Hasher = new(util.MockHasher)
	return authService
}

type InvalidEmailTest struct {
	name       string
	usersInDB  []models.User
	inputEmail string
}

func TestLoginFailsWithInvalidEmail(t *testing.T) {
	tests := []InvalidEmailTest{
		{
			name: "1",
			usersInDB: []models.User{
				{
					Email:    "123@example.com",
					Password: "password",
				},
				{
					Email:    "456@example.com",
					Password: "password",
				},
			},
			inputEmail: "johndoe@example.com",
		},
	}

	for _, test := range tests {
		authService := SetupAuthLoginTest(test.usersInDB)

		token, err := authService.Login(test.inputEmail, "password")

		if want, got := "", token; want != got {
			t.Error("Expected token to be empty")
		}
		if want, got := services.ErrUserNotFound, err; want != got {
			t.Error("Expected error to be of type 'ErrUserNotFound'")
		}
	}
}

func TestLoginFailsWithInvalidPassword(t *testing.T) {
	authService := SetupAuthLoginTest([]models.User{
		{
			Email:    "123@example.com",
			Password: "password",
		},
	})

	token, err := authService.Login("123@example.com", "wrongpassword")

	assert.Empty(t, token)
	assert.Same(t, services.ErrBadPassword, err)
}

func TestLoginSucceedsWithValidInfo(t *testing.T) {
	authService := SetupAuthLoginTest([]models.User{
		{
			Email:    "123@example.com",
			Password: new(util.MockHasher).Hash("password1"),
		},
		{
			Email:    "456@example.com",
			Password: new(util.MockHasher).Hash("password2"),
		},
	})

	token, err := authService.Login("456@example.com", "password2")

	assert.NotEmpty(t, token)
	assert.Nil(t, err)
}
