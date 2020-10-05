package storage

import (
	"errors"

	"github.com/steve-kaufman/react-blog-api/models"
)

// MockUserRepository is a mock repository for users
type MockUserRepository struct {
	Users []models.User
}

// GetByEmail returns a user or an error
func (repo *MockUserRepository) GetByEmail(email string) (models.User, error) {
	for _, user := range repo.Users {
		if user.Email != email {
			continue
		}

		return user, nil
	}

	return models.User{}, errors.New("user not found")
}

// Create creates a user or returns an error
func (repo *MockUserRepository) Create(userData models.User) (models.User, error) {
	return models.User{}, nil
}
