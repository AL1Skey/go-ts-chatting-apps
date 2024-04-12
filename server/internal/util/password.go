package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt's GenerateFromPassword function.
// It returns the hashed password as a string and an error if any occurs during the hashing process.
// The bcrypt.DefaultCost is used as the cost factor for the hashing algorithm.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Error on Hash Password: %w", err)
	}

	return string(hashed), nil
}

// CheckPassword compares the given password with the hashed password.
// It returns an error if the password does not match the hashed password.
func CheckPassword(hashpw string, password string) error {
	// panic(fmt.Sprintf("Password %s, hashed password %s", password, hashpw))
	return bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(password))
}
