package extend

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {

}

// 密码加密
func (b *Bcrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 密码校验
func (b *Bcrypt) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
