package util

import (
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string)string{
	hashpassword,_ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashpassword)
}

func ComparePass(hashpass,pass string)bool{
	if err := bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(pass));
		err != nil {
			return false
	}
	return true

}