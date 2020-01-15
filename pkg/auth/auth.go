package auth

import "golang.org/x/crypto/bcrypt"

/*
 哈希加密
*/
func Encrypt(source string) (string, error)  {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source),bcrypt.DefaultCost)
	return string(hashedBytes), err
}

/*
 校对
*/
func Compare(hashedPasword, password string) error  {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}