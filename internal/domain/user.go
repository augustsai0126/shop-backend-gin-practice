package domain

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	IsAdmin      bool   `gorm:"default:false"`
	Status       string `gorm:"default:'active'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
