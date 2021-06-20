package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
