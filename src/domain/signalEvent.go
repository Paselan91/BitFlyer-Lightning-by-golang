package domain

import (
	"time"
)

type SignalEvent struct {
	Model
	Time        time.Time `json:"time" gorm:"not null"`
	ProductCode string    `json:"productCode" gorm:"not null"`
	Side        string    `json:"side" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Size        float64   `json:"size" gorm:"not null"`
}
