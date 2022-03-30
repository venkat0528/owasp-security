package pwdsecurity

import (
	"golang.org/x/crypto/bcrypt"
)

func GetHashedPassword(password string) (string, error) {
	hashePassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashePassword), err
}
