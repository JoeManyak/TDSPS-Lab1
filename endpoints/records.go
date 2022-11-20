package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"lab1/endpoints/responses"
	le "lab1/local-errors"
	"lab1/models/record"
	"lab1/models/user"
	"log"
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
		responses.Unprocessable(w, record.StructName)
		return
	}

	c, err := record.Parse(data)
	if err != nil {
		responses.UnprocessableDetailed(w, record.StructName, err.Error())
		return
	}

	err = record.Create(c.UserID, c.CategoryID, time.Now(), c.Sum)
	if err != nil {
		if errors.Is(err, le.NotFoundError) {
			responses.NotFound(w, err)
			return
		}
		responses.BadRequest(w, err)
		return
	}
	responses.NoContent(w)
}

func RecordsGet(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(record.GetAll())
	if err != nil {
		responses.Internal(w)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
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

	users := record.GetByUser(int(id))
	data, err := json.Marshal(users)
	if err != nil {
		responses.Internal(w)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
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

	users := record.GetByUserAndCategory(int(userID), int(categoryID))
	if err != nil {
		responses.Unprocessable(w, user.StructName)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		responses.Internal(w)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
