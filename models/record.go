package models

import "time"

var idCount = 0
var records []Record

type Record struct {
	ID         int
	UserID     int
	CategoryID int
	Timestamp  time.Time
	Sum        float64
}

func Create(UserID int, CategoryID int, now time.Time, sum float64) {
	//todo
}
