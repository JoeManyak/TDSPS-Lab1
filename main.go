package main

import (
	"lab1/endpoints"
	"log"
	"net/http"
)

func main() {
	//something like healthcheck
	http.HandleFunc("/", endpoints.Greeting)

	//endpoint to get access to one user
	http.HandleFunc("/user/", endpoints.User)
	//endpoint to get access to list of users
	http.HandleFunc("/users/", endpoints.Users)

	//endpoint to get access to one category
	http.HandleFunc("/category/", endpoints.Category)
	//endpoint to get access to list of categories
	http.HandleFunc("/categories/", endpoints.Categories)

	//endpoint to get access to one record
	http.HandleFunc("/record/", endpoints.Record)
	//endpoint to get access to full list of record
	http.HandleFunc("/records/", endpoints.Records)
	//endpoint to get access to full list of record filtered by user id
	http.HandleFunc("/records/user", endpoints.RecordsByUser)
	//endpoint to get access to full list of record filtered by user id and category id
	http.HandleFunc("/records/user/category", endpoints.RecordsByUserCategory)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
