package endpoints

import (
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Greetings, traveller"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
