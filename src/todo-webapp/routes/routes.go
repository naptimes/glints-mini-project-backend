package routes

import (
	"log"
	"net/http"
	"todo-webapp/handler"

	"github.com/gorilla/mux"
)

func HandleRequest(p string) {
	// available rooutes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.LandingPage)
	router.HandleFunc("/health", handler.Health)
	router.HandleFunc("/register", handler.GetRegister).Methods(http.MethodGet)
	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login).Methods(http.MethodGet)
	router.HandleFunc("/login", handler.LoginPress).Methods(http.MethodPost)
	router.HandleFunc("/login/forgot-password", handler.ForgotPassword)
	router.HandleFunc("/addTodo", handler.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/deleteTodo", handler.DeleteTodo).Methods(http.MethodDelete)

	// err catcher
	err := http.ListenAndServe(":"+p, router)
	log.Fatal(err)
}
