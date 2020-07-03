package util

import (
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string)string{
	hashpassword,_ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashpassword)
}

func Decode(){

}