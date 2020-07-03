package common

import (
	"github.com/fedocx/basicgin/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type Database struct{
	Username string
	Password string
	Port string
	Hostname string
	DBName	string
}

func InitDb()(*gorm.DB,error){

	//Connect := Database{DBName: gindb,Username: root,Password: 123456,Port:3306,Hostname: localhost}
	Connect := GetDBConf()
	connect_string := Connect.Username + ":" + Connect.Password + "@(" + Connect.Hostname + ")/" + Connect.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql",connect_string)
	if err != nil{
		return db,err
	}
	db.AutoMigrate(&model.UserInfo{})
	DB = db
	return db , nil
}

func GetDBConf()*Database{

	Username := viper.GetString("datasource.username")
	Password := viper.GetString("datasource.password")
	Port := viper.GetString("datasource.port")
	Hostname := viper.GetString("datasource.host")
	DBName := viper.GetString("datasource.database")
	db := Database{DBName: DBName,Username: Username,Password: Password,Port: Port,Hostname: Hostname}
	return &db

}

func GetDB() *gorm.DB{
	return DB
}