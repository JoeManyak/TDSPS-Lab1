package main

import (
	"lab1/db"
	ds "lab1/db/db-setup"
	"lab1/endpoints"
	"log"
	"net/http"
	"os"
)

func init() {
	/*	log.Println("Removing latest...")
		if ds.Clear() != nil {
			log.Println("![WARN]! Unable to delete db")
		}
	*/

	log.Println("Connecting to db...")
	con, err := db.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Migrating...")
	err = ds.Migrate(con)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Creating default data...")
	err = ds.CreateDefaultData(con)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

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
	http.HandleFunc("/records/user/", endpoints.RecordsByUser)
	//endpoint to get access to full list of record filtered by user id and category id
	http.HandleFunc("/records/user/category/", endpoints.RecordsByUserCategory)

	log.Printf("Starting server at port: %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
