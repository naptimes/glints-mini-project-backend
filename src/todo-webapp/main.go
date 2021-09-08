package main

import (
	"log"
	"todo-webapp/routes"
)

func main() {
	port := "5000"
	log.Println("Server listion on " + port)
	routes.HandleRequest(port)
}
