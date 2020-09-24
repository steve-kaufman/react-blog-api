package util

// MockHasher is the mock implementation of the PasswordHasher interface
type MockHasher struct{}

// Hash hashes a password
func (MockHasher) Hash(password string) string {
	return "foo" + password + "bar"
}
