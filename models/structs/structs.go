package structs

import (
	"errors"
	"time"
)

// User struct -----
type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Records    []Record   `json:",omitempty" gorm:"foreignKey:UserID"`
	Categories []Category `json:",omitempty" gorm:"foreignKey:CreatedBy"`
}

const UserStructName = "user"

// Record struct -----
type Record struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	CategoryID int
	Timestamp  time.Time
	Sum        float64
}

const RecordStructName = "record"

var ForbiddenCategory = errors.New("cannot use such category")

// Category struct -----
type Category struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedBy int      `gorm:"default:null"`
	Records   []Record `gorm:"foreignKey:CategoryID"`
}

const CategoryStructName = "category"
