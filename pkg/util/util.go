package util

import "golang.org/x/crypto/bcrypt"

func ParsePwd(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}
