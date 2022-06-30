package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Score     uint8  `validate:"gte=0,lte=20"`
	ClassID   uint
	Class     Class `gorm:"references:ID"`
}
