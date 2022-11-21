package endpoints

import (
	"encoding/json"
	"io"
	"lab1/endpoints/responses"
	"lab1/models/category"
	"lab1/models/structs"
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

	var c structs.Category
	err = json.Unmarshal(body, &c)
	if err != nil {
		responses.Unprocessable(w, structs.CategoryStructName)
		return
	}

	cr, err := category.Create(c.Name, c.CreatedBy)
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
