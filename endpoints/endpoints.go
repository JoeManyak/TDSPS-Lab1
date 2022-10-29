package endpoints

import (
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleNotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Greetings, traveller"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func HandleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("Not Found 404, but you can give me 60! :)"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
