package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name  string    `json:"name"`
	Brand string    `json:"brand"`
	Year  time.Time `json:"year"`
}
