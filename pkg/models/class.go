package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Field   string
	Teacher string
}
