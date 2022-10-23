package endpoints

import (
	"encoding/json"
	"io"
	"lab1/endpoints/responses"
	"lab1/models"
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
	}
}

// Records represents endpoints from route /records/
func Records(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGet(w, r)
		break
	}
}

// RecordsByUser represents endpoints from route /records/user/
func RecordsByUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGetByUser(w, r)
		break
	}
}

// RecordsByUserCategory represents endpoints from route /records/user/category
func RecordsByUserCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RecordsGetByUserCategory(w, r)
		break
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
		responses.Unprocessable(w, models.StructName)
		return
	}

	c, err := models.Parse(data)
	if err != nil {
		responses.UnprocessableDetailed(w, models.StructName, err.Error())
		return
	}

	err = models.Create(c.UserID, c.CategoryID, time.Now(), c.Sum)
	if err != nil {
		responses.BadRequest(w, err)
		return
	}
	responses.NoContent(w)
}

func RecordsGet(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(models.GetAll())
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

	id, ok := request["user_id"].(int)
	if !ok {
		responses.Unprocessable(w, "user_id")
		return
	}

	users := models.GetByUser(id)
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

	userID, ok := request["user_id"].(int)
	if !ok {
		responses.Unprocessable(w, "user_id")
		return
	}

	categoryID, ok := request["category_id"].(int)
	if !ok {
		responses.Unprocessable(w, "category_id")
		return
	}

	users := models.GetByUserAndCategory(userID, categoryID)
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
