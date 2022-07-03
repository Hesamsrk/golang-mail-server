package handlers

import (
	"fmt"
	"net/http"

	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"github.com/labstack/echo/v4"
)

func (h withDBHandler) SaveStudent(ctx echo.Context) error {
	var student = &models.Student{}
	if err := ctx.Bind(&student); err != nil {
		return err
	}
	if err := ctx.Validate(student); err != nil {
		return err
	}

	var found models.Student
	h.DB.Model(&found).Where("email = ?", student.Email).First(&found)

	if student.Email == found.Email {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"massage": "Student with this email already exists",
		})
	}

	if student.ClassID <= 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"massage": "No 'classID' provided.",
		})
	}
	var class = &models.Class{}
	h.DB.First(&class, student.ClassID)

	fmt.Printf("******\n%v\n******", *class)

	if class.ID == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"massage": "No class found with the 'classID' you proovided.",
		})
	}

	h.DB.Create(&student)

	if student.ID == 0 {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"massage": "No student created",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"ID": student.ID,
	})
}

func (h withDBHandler) RemoveStudent(ctx echo.Context) error {
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
	student := &models.Student{}

	h.DB.Delete(&student, req.ID)

	return ctx.JSON(http.StatusOK, struct {
		Status string
	}{
		Status: "OK",
	})

}

func (h withDBHandler) GetStudntsList(ctx echo.Context) error {
	var students []models.Student
	result := h.DB.Joins("Class").Find(&students)
	if result.Error != nil {
		return result.Error
	}
	m := make(map[int]interface{})

	for index, user := range students {
		m[index] = user
	}

	return ctx.JSON(http.StatusOK, m)
}
