package storage

import "github.com/steve-kaufman/react-blog-api/models"

// UserRepository has methods for interacting with a storage repo
type UserRepository interface {
	Create(userData models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
}
