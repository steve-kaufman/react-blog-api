package util

// DefaultHasher is the default object for hashing passwords
type DefaultHasher struct{}

// Hash hashes a password
func (DefaultHasher) Hash(string) string {
	return "default"
}
