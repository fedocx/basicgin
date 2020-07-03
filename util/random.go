package util

import (
	"math/rand"
	"time"
)


func CreateUserName(num int)string{
	return RandomString(num)
}

func RandomString(num int)string{
	var letter = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, num)
	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}