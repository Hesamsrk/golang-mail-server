package handlers

import (
	"net/http"

	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h withDBHandler) GetUserList(ctx echo.Context) error {
	var users []models.User
	result := h.DB.Find(&users)
	if result.Error != nil {
		return result.Error
	}
	m := make(map[int]interface{})

	for index, user := range users {
		m[index] = user
	}

	return ctx.JSON(http.StatusOK, m)
}
