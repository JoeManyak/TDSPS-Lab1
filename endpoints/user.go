package endpoints

import (
	"encoding/json"
	"io"
	"lab1/endpoints/responses"
	"lab1/models/structs"
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
	default:
		HandleNotFound(w, r)
	}
}

// Users represents endpoints from route /users/
func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		UsersGet(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var u structs.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		responses.Unprocessable(w, structs.UserStructName)
		return
	}

	u, err = user.Create(u.Name)
	if err != nil {
		responses.BadRequest(w, err)
		return
	}

	responses.OK(w, u)
}

func UsersGet(w http.ResponseWriter, _ *http.Request) {
	all, err := user.GetAll()
	if err != nil {
		responses.Internal(w)
		return
	}

	data, err := json.Marshal(all)
	if err != nil {
		responses.Internal(w)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
