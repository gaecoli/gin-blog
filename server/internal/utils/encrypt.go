package utils

import "golang.org/x/crypto/bcrypt"

func CryptCheck(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
