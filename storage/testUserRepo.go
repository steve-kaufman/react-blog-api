package storage

import (
	"errors"

	"github.com/steve-kaufman/react-blog-api/models"
)

// TestUserRepository is a mock repository for users
type TestUserRepository struct {
	Users []models.User
}

// GetByEmail returns a user or an error
func (repo *TestUserRepository) GetByEmail(email string) (models.User, error) {
	for _, user := range repo.Users {
		if user.Email != email {
			continue
		}

		return user, nil
	}

	return models.User{}, errors.New("user not found")
}