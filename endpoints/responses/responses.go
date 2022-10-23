package responses

import (
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
