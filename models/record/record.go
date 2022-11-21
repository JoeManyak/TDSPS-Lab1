package record

import (
	"fmt"
	"gorm.io/gorm"
	"lab1/db"
	le "lab1/local-errors"
	"lab1/models/structs"
	"time"
)

func Parse(data map[string]interface{}) (structs.Record, error) {
	var result structs.Record
	if v, ok := data["user_id"].(float64); ok {
		result.UserID = int(v)
	} else {
		return structs.Record{}, le.NoField("user_id")
	}
	if v, ok := data["category_id"].(float64); ok {
		result.CategoryID = int(v)
	} else {
		return structs.Record{}, le.NoField("category_id")
	}
	if v, ok := data["sum"].(float64); ok {
		result.Sum = v
	} else {
		return structs.Record{}, le.NoField("sum")
	}

	return result, nil
}

func Create(userID int, categoryID int, now time.Time, sum float64) (structs.Record, error) {
	connect, err := db.Connect()
	if err != nil {
		return structs.Record{}, fmt.Errorf("create record on connect: %w", err)
	}

	if !checkCategory(connect, userID, categoryID) {
		return structs.Record{}, structs.ForbiddenCategory
	}

	r := structs.Record{
		UserID:     userID,
		CategoryID: categoryID,
		Timestamp:  now,
		Sum:        sum,
	}

	tx := connect.Create(&r)
	if tx.Error != nil {
		return structs.Record{}, tx.Error
	}

	return r, nil
}

func GetAll() ([]structs.Record, error) {
	var result []structs.Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records on connect: %w", err)
	}

	tx := connect.Model(structs.Record{}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records: %w", err)
	}

	return result, nil
}

func GetByUser(id int) ([]structs.Record, error) {
	var result []structs.Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records by user on connect: %w", err)
	}

	tx := connect.Model(structs.Record{}).Where(&structs.Record{UserID: id}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records by user: %w", err)
	}

	return result, nil
}

func GetByUserAndCategory(userID, categoryID int) ([]structs.Record, error) {
	var result []structs.Record

	connect, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("get records by user on connect: %w", err)
	}

	tx := connect.Model(structs.Record{}).Where(&structs.Record{UserID: userID, CategoryID: categoryID}).Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("get records by user: %w", err)
	}

	return result, nil
}

func checkCategory(connect *gorm.DB, userID, categoryID int) bool {
	var c structs.Category

	tx := connect.
		Model(&structs.Category{}).
		Where("id = ? AND (created_by = ? OR created_by = NULL)", categoryID, userID).
		Find(&c)

	if c.ID != categoryID || tx.Error != nil {
		return false
	}

	return true
}
