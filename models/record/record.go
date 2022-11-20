package record

import (
	"fmt"
	"lab1/db"
	le "lab1/local-errors"
	"lab1/models/category"
	"time"
)

const StructName = "record"

type Record struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	CategoryID int
	Timestamp  time.Time
	Sum        float64
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

func Create(userID int, categoryID int, now time.Time, sum float64) (Record, error) {
	connect, err := db.Connect()
	if err != nil {
		return Record{}, fmt.Errorf("create record on connect: %w", err)
	}

	if !category.CheckCategory(connect, userID, categoryID) {
		return Record{}, category.ForbiddenCategory
	}

	r := Record{
		UserID:     userID,
		CategoryID: categoryID,
		Timestamp:  now,
		Sum:        sum,
	}

	tx := connect.Create(&r)
	if tx.Error != nil {
		return Record{}, tx.Error
	}

	return r, nil
}

func GetAll() ([]Record, error) {
	var result []Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records on connect: %w", err)
	}

	tx := connect.Model(Record{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records: %w", err)
	}

	return result, nil
}

func GetByUser(id int) ([]Record, error) {
	var result []Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records by user on connect: %w", err)
	}

	tx := connect.Model(Record{}).Where(&Record{UserID: id}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records by user: %w", err)
	}

	return result, nil
}

func GetByUserAndCategory(userID, categoryID int) ([]Record, error) {
	var result []Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records by user on connect: %w", err)
	}

	tx := connect.Model(Record{}).Where(&Record{UserID: userID, CategoryID: categoryID}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records by user: %w", err)
	}

	return result, nil
}
