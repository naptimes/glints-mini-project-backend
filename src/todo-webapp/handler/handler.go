package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dns = "todoAdmin:IniPassword456$@tcp(127.0.0.1:3306)/todowebapp?charset=utf8mb4&parseTime=True&loc=Local"

var DB, Err = gorm.Open(mysql.Open(dns), &gorm.Config{})

// type User struct {
// 	FirstName     string   `json:"fname"`
// 	Lastname      string   `json:"lname"`
// 	LoginAttribut LoginAtt `json:"login-att"`
// }

type User struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAtt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Task struct {
	TaskID     int    `json:"taskId"`
	Item       string `json:"item"`
	TaskStatus bool   `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// FIX
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var user User
	json.Unmarshal(body, &user)

	var res Result

	if user.Email == "" || user.Password == "" || user.Fname == "" || user.Lname == "" {
		res = Result{Status: 406, Message: http.StatusText(http.StatusNotAcceptable)}
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		if dbe := DB.Create(&user); dbe.Error != nil {
			res = Result{Status: 406, Message: http.StatusText(http.StatusNotAcceptable)}
			w.WriteHeader(http.StatusNotAcceptable)
		} else {
			res = Result{Status: 201, Message: http.StatusText(http.StatusAccepted)}
			w.WriteHeader(http.StatusAccepted)
		}
	}

	result, _ := json.Marshal(res)

	w.Write(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route login"))
}

// FIX
func LoginPress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var user User
	var catch User
	var res Result

	json.Unmarshal(body, &catch)

	DB.First(&user, catch)
	if user.Email == "" {
		res = Result{Status: 406, Message: http.StatusText(http.StatusNotAcceptable)}
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		res = Result{Status: 200, Message: http.StatusText(http.StatusOK)}
		w.WriteHeader(http.StatusOK)
	}

	result, _ := json.Marshal(res)

	w.Write(result)
}

func GetRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route GET register/"))
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route login/forgot-password"))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini landing page~"))
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var task Task
	json.Unmarshal(body, &task)

	var res Result

	if task.Item == "" {
		res = Result{Status: 406, Message: http.StatusText(http.StatusNotAcceptable)}
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		if dbe := DB.Create(&task); dbe.Error != nil {
			res = Result{Status: 406, Message: http.StatusText(http.StatusNotAcceptable)}
			w.WriteHeader(http.StatusNotAcceptable)
		} else {
			res = Result{Status: 201, Message: http.StatusText(http.StatusAccepted)}
			w.WriteHeader(http.StatusAccepted)
		}
	}

	result, _ := json.Marshal(res)

	w.Write(result)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var task Task
	var catch Task
	var res Result

	json.Unmarshal(body, &catch)

	DB.Delete(&task, "task_id = ?", catch.TaskID)
	res = Result{Status: 200, Message: http.StatusText(http.StatusOK)}
	w.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(res)

	w.Write(result)
}
