package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:root@/sms?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db.LogMode(true)
		return
	}
	panic(err.Error())
}
