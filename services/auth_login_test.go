package services_test

import (
	"testing"

	"github.com/steve-kaufman/react-blog-api/models"
	"github.com/steve-kaufman/react-blog-api/services"
	"github.com/steve-kaufman/react-blog-api/storage"
	"github.com/steve-kaufman/react-blog-api/util"
)

func SetupAuthLoginTest() *services.AuthService {
	mockHasher := new(util.MockHasher)

	repo := new(storage.MockUserRepository)
	repo.Users = []models.User{
		{
			Email:    "123@example.com",
			Password: mockHasher.Hash("password1"),
		},
		{
			Email:    "456@website.com",
			Password: mockHasher.Hash("password2"),
		},
		{
			Email:    "789@foo.com",
			Password: mockHasher.Hash("password3"),
		},
		{
			Email:    "johndoe@bar.com",
			Password: mockHasher.Hash("password4"),
		},
		{
			Email:    "johndoe@example.com",
			Password: mockHasher.Hash("password5"),
		},
		{
			Email:    "janedoe@example.com",
			Password: mockHasher.Hash("password6"),
		},
	}

	authService := services.NewAuthService(repo)
	authService.Hasher = new(util.MockHasher)
	return authService
}

type BadLoginTest struct {
	name          string
	inputEmail    string
	inputPassword string
	expectedError error
}

func TestLoginFailsWithInvalidCredentials(t *testing.T) {
	tests := []BadLoginTest{
		{
			name:          "wrong name right domain right password",
			inputEmail:    "nobody@example.com",
			inputPassword: "password1",
			expectedError: services.ErrUserNotFound,
		},
		{
			name:          "wrong name right domain wrong password",
			inputEmail:    "nobody@example.com",
			inputPassword: "badpassword",
			expectedError: services.ErrUserNotFound,
		},
		{
			name:          "wrong domain right name right password",
			inputEmail:    "123@wrong.com",
			inputPassword: "password1",
			expectedError: services.ErrUserNotFound,
		},
		{
			name:          "wrong name wrong domain",
			inputEmail:    "bad@wrong.com",
			inputPassword: "password1",
			expectedError: services.ErrUserNotFound,
		},
		{
			name:          "right name right domain wrong password",
			inputEmail:    "johndoe@example.com",
			inputPassword: "badpassword",
			expectedError: services.ErrBadPassword,
		},
		{
			name:          "right name right domain wrong password",
			inputEmail:    "789@foo.com",
			inputPassword: "badpassword",
			expectedError: services.ErrBadPassword,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authService := SetupAuthLoginTest()

			token, err := authService.Login(test.inputEmail, test.inputPassword)

			if want, got := "", token; want != got {
				t.Error("Expected token to be empty")
			}
			if want, got := test.expectedError, err; want != got {
				t.Error("Expected error to be of type 'ErrUserNotFound'")
			}
		})
	}
}

type GoodLoginTest struct {
	name          string
	inputEmail    string
	inputPassword string
}

func TestLoginSucceedsWithValidInfo(t *testing.T) {
	authService := SetupAuthLoginTest()

	tests := []GoodLoginTest{
		{
			name:          "1",
			inputEmail:    "123@example.com",
			inputPassword: "password1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token, err := authService.Login(test.inputEmail, test.inputPassword)

			if token == "" {
				t.Error("Expected a token")
			}
			if err != nil {
				t.Error("Expected no error")
			}
		})
	}
}
