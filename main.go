package main

import (
	"log"
	"net/http"

	"example.com/appointyTask/handlers"
)

func main() {
	http.HandleFunc("/users/", handlers.GetUserHandler)
	http.HandleFunc("/contacts", handlers.ContactHandler)
	log.Println("starting server")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
