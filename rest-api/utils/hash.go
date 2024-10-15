package utils

import "golang.org/x/crypto/bcrypt"

// Generate password as hashed format

func GenerateHashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(hashedPassword)
}

func CompareHashPassword(hashedPassword, password string) bool {
	// Compare the hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
