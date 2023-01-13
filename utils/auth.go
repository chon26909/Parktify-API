package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {

	salt := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
