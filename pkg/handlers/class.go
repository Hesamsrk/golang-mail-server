package handlers

import (
	"net/http"

	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h withDBHandler) SaveClass(ctx echo.Context) error {
	var class = &models.Class{}
	if err := ctx.Bind(&class); err != nil {
		return err
	}
	if err := ctx.Validate(class); err != nil {
		return err
	}

	h.DB.Create(&class)

	if class.ID == 0 {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"massage": "No class created",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"ID": class.ID,
	})
}

func (h withDBHandler) RemoveClass(ctx echo.Context) error {
	type request struct {
		ID int `json:"id"`
	}
	req := &request{}

	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return err
	}
	class := &models.Class{}
	h.DB.Delete(&class, req.ID)

	return ctx.JSON(http.StatusOK, struct {
		Status string
	}{
		Status: "OK",
	})

}

func (h withDBHandler) GetClassesList(ctx echo.Context) error {
	var classes []models.Class
	result := h.DB.Find(&classes)

	if result.Error != nil {
		return result.Error
	}
	m := make(map[int]interface{})

	for index, user := range classes {
		m[index] = user
	}

	return ctx.JSON(http.StatusOK, m)
}
