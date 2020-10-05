package storage

import (
	"github.com/steve-kaufman/react-blog-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User is the GORM model for users
type User struct {
	gorm.Model
	models.User
}

// SqliteUserRepository is the sqlite implementation of UserRepository
type SqliteUserRepository struct {
	db *gorm.DB
}

// GetByEmail gets a user by email or returns an error
func (repo SqliteUserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	repo.db.First(&user, models.User{Email: email})

	return user, nil
}

// Create creates a user or returns an error
func (repo SqliteUserRepository) Create(userData models.User) (models.User, error) {
	user := User{gorm.Model{}, userData}

	repo.db.Create(&user)

	return models.User{}, nil
}

// NewSqliteUserRepository creates a database
func NewSqliteUserRepository() *SqliteUserRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		panic("failed to connect to database")
	}

	repo := new(SqliteUserRepository)
	repo.db = db

	return repo
}
