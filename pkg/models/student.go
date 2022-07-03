package models

type Student struct {
	Model
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `gorm:"unique" validate:"required,email"`
	Score     uint8  `validate:"required,gte=0,lte=20"`
	ClassID   uint
	Class     Class `gorm:"references:ID"`
}
