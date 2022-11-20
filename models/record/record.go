package record

import (
	le "lab1/local-errors"
	"time"
)

var idCount = 0
var records []Record

const StructName = "record"

type Record struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	CategoryID int
	Timestamp  time.Time
	Sum        float64
}

func init() {
	records = make([]Record, 0, 10)
}

func Parse(data map[string]interface{}) (Record, error) {
	var result Record
	if v, ok := data["user_id"].(float64); ok {
		result.UserID = int(v)
	} else {
		return Record{}, le.NoField("user_id")
	}
	if v, ok := data["category_id"].(float64); ok {
		result.CategoryID = int(v)
	} else {
		return Record{}, le.NoField("category_id")
	}
	if v, ok := data["sum"].(float64); ok {
		result.Sum = v
	} else {
		return Record{}, le.NoField("sum")
	}

	return result, nil
}

func Create(userID int, categoryID int, now time.Time, sum float64) error {
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

func GetAll() []Record {
	return records
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
