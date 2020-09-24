package util

// PasswordHasher is the function type for a password hasher
type PasswordHasher interface {
	Hash(password string) (hash string)
}
