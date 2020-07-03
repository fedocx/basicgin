package model

import "github.com/jinzhu/gorm"

type UserInfo struct{
	gorm.Model
	Username string `gorm:"varchar(30),not null"`
	Password string `gorm:"size:255:not null"`
	Telephone string	`gorm:"varchar(11);not null;unique"`
}
