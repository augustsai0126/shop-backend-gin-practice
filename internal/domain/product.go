package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	ImageURL    string
	Status      string `gorm:"default:active"`
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
