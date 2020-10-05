package storage_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/steve-kaufman/react-blog-api/models"
	"github.com/steve-kaufman/react-blog-api/storage"
)

func SetupSqliteUsersTest() *storage.SqliteUserRepository {
	os.Setenv("GO_API_SQLITE_PATH", "./test.db")

	os.Remove("./test.db")

	return storage.NewSqliteUserRepository()
}

func TestSqliteUsers_CreatesDatabase(t *testing.T) {
	SetupSqliteUsersTest() // deletes test.db

	_, err := os.Open("./test.db")

	if err != nil {
		t.Error("Expected db to be created")
	}
}

func TestSqliteUsers_AddsUsersTable(t *testing.T) {
	SetupSqliteUsersTest()

	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if !db.Migrator().HasTable("users") {
		t.Error("Expected users table to be created")
	}
}

func TestSqliteUsers_AddsUser(t *testing.T) {
	usersStorage := SetupSqliteUsersTest()

	usersStorage.Create(models.User{
		Email:    "123@example.com",
		Password: "supersecret",
	})

	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	var user models.User

	db.First(&user, models.User{Email: "123@example.com"})

	if (user == models.User{}) {
		t.Error("Was expecting user to exist")
	}

	if user.Password != "supersecret" {
		t.Error("Was expecting passwords to match")
	}
}

type GetByEmailTest struct {
	name              string
	usersInDatabase   []models.User
	inputEmail        string
	expectedUserIndex int
}

func (test GetByEmailTest) expectedUser() models.User {
	return test.usersInDatabase[test.expectedUserIndex]
}

func TestSqliteUsers_GetsUserByEmail(t *testing.T) {
	tests := []GetByEmailTest{
		{
			name: "first",
			usersInDatabase: []models.User{
				{
					ID:       1,
					Email:    "123@example.com",
					Password: "pass1",
				},
				{
					ID:       2,
					Email:    "456@example.com",
					Password: "pass2",
				},
				{
					ID:       3,
					Email:    "789@example.com",
					Password: "pass3",
				},
			},
			inputEmail:        "123@example.com",
			expectedUserIndex: 0,
		},
		{
			name: "second",
			usersInDatabase: []models.User{
				{
					ID:       1,
					Email:    "123@example.com",
					Password: "pass1",
				},
				{
					ID:       2,
					Email:    "456@example.com",
					Password: "pass2",
				},
				{
					ID:       3,
					Email:    "789@example.com",
					Password: "pass3",
				},
			},
			inputEmail:        "456@example.com",
			expectedUserIndex: 1,
		},
		{
			name: "third",
			usersInDatabase: []models.User{
				{
					ID:       1,
					Email:    "123@example.com",
					Password: "pass1",
				},
				{
					ID:       2,
					Email:    "456@example.com",
					Password: "pass2",
				},
				{
					ID:       3,
					Email:    "789@example.com",
					Password: "pass3",
				},
			},
			inputEmail:        "789@example.com",
			expectedUserIndex: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usersStorage := SetupSqliteUsersTest()

			db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
			db.Create(&test.usersInDatabase)

			user, err := usersStorage.GetByEmail(test.inputEmail)

			if err != nil {
				t.Error("Expected no error")
			}

			if want, got := user, test.expectedUser(); cmp.Diff(want, got) != "" {
				t.Error("Expected user to match")
			}
		})
	}
}
