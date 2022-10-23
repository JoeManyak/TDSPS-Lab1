package endpoints

import (
	"encoding/json"
	"io"
	"lab1/endpoints/responses"
	"lab1/models/user"
	"log"
	"net/http"
)

// User represents endpoints from route /user/
func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		UserCreate(w, r)
		break
	}
}

// Users represents endpoints from route /users/
func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UserGet(w, r)
		break
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var u user.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		responses.Unprocessable(w, user.StructName)
		return
	}

	user.Create(u.Name)
	responses.NoContent(w)
}

func UserGet(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(user.GetAll())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
