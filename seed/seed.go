package main

import (
	"example/goRestAPI/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_orm_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Username: "aom31", Fullname: "thamakorn", Password: "aom1234", Email: "mamjia@gmail.com"})
	db.Create(&model.User{Username: "a1345", Fullname: "kornaoka", Password: "1234", Email: "thamakorn@gmail.com"})

}
