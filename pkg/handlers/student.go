package handlers

import (
	"github.com/Hesamsrk/golang-mail-server/pkg/models"
)

func (h withDBHandler) CreateStudent(input models.Student) {
	h.DB.Create(&input)
}
