package domain

import (
	"time"
)

type BtcCandle struct {
	ProductCode string        `json:"productCode" gorm:"not null"`
	Duration    time.Duration `json:"duration" gorm:"not null"`
	Time        time.Time     `json:"time" gorm:"not null"`
	Open        float64       `json:"open" gorm:"not null"`
	Close       float64       `json:"close" gorm:"not null"`
	High        float64       `json:"high" gorm:"not null"`
	Low         float64       `json:"low" gorm:"not null"`
	Volume      float64       `json:"volume" gorm:"not null"`
}

func NewCandle(productCode string, duration time.Duration, timeDate time.Time, open, close, high, low, volume float64) *BtcCandle {
	return &BtcCandle{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

type Btc1sCandle struct {
	Model
	BtcCandle
}
type Btc1mCandle struct {
	Model
	BtcCandle
}
type Btc1hCandle struct {
	Model
	BtcCandle
}
