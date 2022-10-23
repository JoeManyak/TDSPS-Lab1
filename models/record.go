package models

import (
	"lab1/models/category"
	"lab1/models/user"
	"time"
)

var idCount = 0
var records []Record

type Record struct {
	ID         int
	UserID     int
	CategoryID int
	Timestamp  time.Time
	Sum        float64
}

func Create(userID int, categoryID int, now time.Time, sum float64) error {
	_, err := user.GetByID(userID)
	if err != nil {
		return err
	}

	_, err = category.GetByID(categoryID)
	if err != nil {
		return err
	}

	records = append(records, Record{
		ID:         idCount,
		UserID:     userID,
		CategoryID: categoryID,
		Timestamp:  now,
		Sum:        sum,
	})
	idCount++
	return nil
}

func GetByUser(id int) []Record {
	var result []Record
	for i := range records {
		if records[i].UserID == id {
			result = append(result, records[i])
		}
	}
	return result
}

func GetByUserAndCategory(userID, categoryID int) []Record {
	var result []Record
	for i := range records {
		if records[i].UserID == userID && records[i].CategoryID == categoryID {
			result = append(result, records[i])
		}
	}
	return result
}