package responses

import (
	"fmt"
	"log"
	"net/http"
)

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func Internal(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func Unprocessable(w http.ResponseWriter, structName string) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(fmt.Sprintf("provided %s is invalid", structName)))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
