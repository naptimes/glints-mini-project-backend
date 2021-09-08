package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	dsn := "todoAdmin:IniPassword456$@tcp(127.0.0.1:3306)/todowebapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
	} else {
		log.Println("Connection established")
	}

	// CREATE
	user := User{Fname: "test", Lname: "name", Email: "mock@email.com", Password: "IniMockPass"}
	db.Create(&user)
	// trying duplicate email
	db.Create(&user)

	// READ
	// var getUser User
	// db.First(&getUser, "fname = ?", "test") // find product with code D42
	// fmt.Println(getUser)

	// Delete - delete product
	// db.Delete(&user, 1)
}
