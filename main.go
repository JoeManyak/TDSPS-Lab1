package main

import (
	"lab1/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoints.Greeting)

	//endpoint to get access to one user
	http.HandleFunc("/user/", endpoints.User)

	//endpoint to get access to list of users
	http.HandleFunc("/users/", endpoints.Users)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
