package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	FName    string `json:"firstname", gorm:"not null"`
	LName    string `json:"lastname"`
	Email    string `json:"email", gorm:"not null;unique"`
	Password string `json:"password", gorm:"not null"`
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic("Could not connect to the database")
	}
	defer db.Close()

	/*var t User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&t)
	w.Write([]byte(t.Email))
	email2 := db.Model(&User{}).Select("email").Where("name LIKE ?", t.Email)
	w.Write([]byte(email2))*/

	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)
	json.NewEncoder(w).Encode(user)
	/*vars := mux.Vars(r)
	fname := vars["fname"]
	lname := vars["lname"]
	email := vars["email"]
	password := vars["password"]*/

	//db.Create(&User{FName: fname, LName: lname, Email: email, Password: password})

	fmt.Fprintf(w, "New User Successfully Created")
}
