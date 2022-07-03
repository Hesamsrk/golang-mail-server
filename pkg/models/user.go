package models


type User struct {
	Model
	Username string `validate:"required,min=5" gorm:"unique"`
	Password string `validate:"required,min=5"`
	Token    string `validate:"omitempty"`
}
