package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName     string   `json:"fname"`
	Lastname      string   `json:"lname"`
	LoginAttribut LoginAtt `json:"login-att"`
}

type LoginAtt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route register"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route login"))
}

func LoginPress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	loginTest := LoginAtt{
		Email: "mockb@mal.com", Password: "IniMockPassword123@"}

	// for checking user credentail (temporary)
	if loginTest.Email != "mockb@mail.com" || loginTest.Password != "IniMockPassword123@" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		w.Write([]byte("terjadi error"))
		return
	}
	json.NewEncoder(w).Encode(loginTest)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini route login/forgot-password"))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini landing page~"))
}
