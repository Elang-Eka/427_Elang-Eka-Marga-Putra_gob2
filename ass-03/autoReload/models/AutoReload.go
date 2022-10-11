package models

import "github.com/jinzhu/gorm"

// model/struct Order
type AutoReload struct {
	gorm.Model
	Water int `json:"water" `
	Wind  int `json:"wind"`
}
