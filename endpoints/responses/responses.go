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

//func NotFound(w http.ResponseWriter, structName string) {
//	w.WriteHeader(http.StatusNotFound)
//	_, err := w.Write([]byte(fmt.Sprintf("%s not found", structName)))
//	if err != nil {
//		log.Fatalln(err.Error())
//		return
//	}
//}

func BadRequest(w http.ResponseWriter, errorStr error) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(errorStr.Error()))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
