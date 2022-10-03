package models

import "time"

// Declaring Models(tags)
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdateAt  time.Time
}
