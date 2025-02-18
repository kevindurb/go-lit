package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Description string `json:"description"`
}
