package handlers

import (
	"github.com/Hesamsrk/golang-mail-server/pkg/models"
)

func (h withDBHandler) CreateClass(input models.Class) {
	h.DB.Create(&input)
}
