package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"lab1/endpoints/responses"
	"lab1/models/category"
	"lab1/models/structs"
	"lab1/models/user"
	"net/http"
)

// Category represents endpoints from route /category/
func Category(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CategoryCreate(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

// Categories represents endpoints from route /categories/
func Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		CategoriesGet(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var raw map[string]any
	err = json.Unmarshal(body, &raw)
	if err != nil {
		responses.Unprocessable(w, structs.CategoryStructName)
		return
	}

	parsed, ok := category.Parse(raw)
	if !ok {
		responses.Unprocessable(w, structs.CategoryStructName)
		return
	}

	u, err := user.GetByID(parsed.CreatedBy)
	if err != nil {
		responses.BadRequest(w, err)
		return
	}

	if user.Equal(u, structs.User{}) {
		responses.NotFound(w, errors.New("no such user"))
		return
	}

	if parsed.Name == "" {
		responses.BadRequest(w, errors.New("name cannot be empty"))
		return
	}

	cr, err := category.Create(parsed.Name, parsed.CreatedBy)
	if err != nil {
		return
	}

	responses.OK(w, cr)
}

func CategoriesGet(w http.ResponseWriter, _ *http.Request) {
	categories, err := category.GetAll()
	if err != nil {
		responses.Internal(w)
		return
	}

	responses.OK(w, categories)
}
