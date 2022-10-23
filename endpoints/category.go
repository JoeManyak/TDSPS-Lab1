package endpoints

import (
	"encoding/json"
	"io"
	"lab1/endpoints/responses"
	"lab1/models/category"
	"lab1/models/user"
	"log"
	"net/http"
)

// Category represents endpoints from route /category/
func Category(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CategoryCreate(w, r)
		break
	}
}

// Categories represents endpoints from route /categories/
func Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		CategoriesGet(w, r)
		break
	}
}

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var c category.Category
	err = json.Unmarshal(body, &c)
	if err != nil {
		responses.Unprocessable(w, user.StructName)
		return
	}

	category.Create(c.Name)
	responses.NoContent(w)
}

func CategoriesGet(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(category.GetAll())
	if err != nil {
		responses.Internal(w)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
