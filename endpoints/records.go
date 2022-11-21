package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"lab1/endpoints/responses"
	le "lab1/local-errors"
	rec "lab1/models/record"
	"lab1/models/structs"
	"net/http"
	"time"
)

// Record represents endpoints from route /record/
func Record(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		RecordCreate(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

// Records represents endpoints from route /records/
func Records(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGet(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

// RecordsByUser represents endpoints from route /records/user/
func RecordsByUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGetByUser(w, r)
		break
	default:
		HandleNotFound(w, r)
	}

}

// RecordsByUserCategory represents endpoints from route /records/user/category
func RecordsByUserCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGetByUserCategory(w, r)
		break
	default:
		HandleNotFound(w, r)
	}
}

func RecordCreate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.Unprocessable(w, structs.RecordStructName)
		return
	}

	c, err := rec.Parse(data)
	if err != nil {
		responses.UnprocessableDetailed(w, structs.RecordStructName, err.Error())
		return
	}

	re, err := rec.Create(c.UserID, c.CategoryID, time.Now(), c.Sum)
	if err != nil {
		if errors.Is(err, structs.ForbiddenCategory) {
			responses.Forbidden(w, err)
			return
		}
		if errors.Is(err, le.NotFoundError) {
			responses.NotFound(w, err)
			return
		}

		responses.BadRequest(w, err)
		return
	}

	responses.OK(w, re)
}

func RecordsGet(w http.ResponseWriter, _ *http.Request) {
	records, err := rec.GetAll()
	if err != nil {
		responses.Internal(w)
		return
	}

	responses.OK(w, records)
}

func RecordsGetByUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var request map[string]interface{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		responses.Unprocessable(w, "map")
		return
	}

	id, ok := request["user_id"].(float64)
	if !ok {
		responses.Unprocessable(w, "user_id")
		return
	}

	users, err := rec.GetByUser(int(id))
	if err != nil {
		responses.NotFound(w, err)
		return
	}

	responses.OK(w, users)
}

func RecordsGetByUserCategory(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Internal(w)
		return
	}

	var request map[string]interface{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		responses.Unprocessable(w, "map")
		return
	}

	userID, ok := request["user_id"].(float64)
	if !ok {
		responses.Unprocessable(w, "user_id")
		return
	}

	categoryID, ok := request["category_id"].(float64)
	if !ok {
		responses.Unprocessable(w, "category_id")
		return
	}

	users, err := rec.GetByUserAndCategory(int(userID), int(categoryID))
	if err != nil {
		responses.NotFound(w, err)
		return
	}

	responses.OK(w, users)
}
