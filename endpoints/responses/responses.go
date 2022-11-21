package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func OK(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		Internal(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func Internal(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(""))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func Unprocessable(w http.ResponseWriter, structName string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, err := w.Write([]byte(fmt.Sprintf("provided %s is invalid", structName)))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func UnprocessableDetailed(w http.ResponseWriter, structName string, additional string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, err := w.Write([]byte(fmt.Sprintf("provided %s is invalid: %s", structName, additional)))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func NotFound(w http.ResponseWriter, errGet error) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(errGet.Error()))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func Forbidden(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusForbidden)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func BadRequest(w http.ResponseWriter, errorStr error) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte(errorStr.Error()))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
