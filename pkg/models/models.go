package models

import (
	"time"
)

type Position struct {
	Ticker    string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Type      string
	Count     uint16
	Price     uint32
	Amount    uint32
}
