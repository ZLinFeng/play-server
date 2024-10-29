package common

import "golang.org/x/crypto/bcrypt"

func BcryptHash(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash)
}
