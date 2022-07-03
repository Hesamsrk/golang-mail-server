package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"github.com/Hesamsrk/golang-mail-server/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (h withDBHandler) SendClassDataToStudents(ctx echo.Context) error {
	type request struct {
		ClassID int `json:"class_id"`
	}
	req := &request{}

	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return err
	}

	//Find class:
	var class models.Class
	result := h.DB.First(&class, req.ClassID)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"massage": "No class found with this ID!",
		})
	}
	// Find students:
	var students []models.Student
	result = h.DB.Select("first_name", "last_name", "email", "score").Where("class_id = ?", req.ClassID).Find(&students)
	if result.Error != nil || len(students) == 0 {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"massage": "Student found  in this class!",
		})
	}

	var wg sync.WaitGroup
	var errors []error
	for _, student := range students {
		wg.Add(1)
		mailText := fmt.Sprintf("Hello Dear Mr./Ms. %s %s\nYour score in <b>%s</b> class by Professor <b>%s</b> is %d .", student.FirstName, student.LastName, class.Field, class.Teacher, student.Score)
		go func() {
			defer wg.Done()
			err := utils.SendMail(student.Email, mailText)
			if err != nil {
				errors = append(errors, err)
			}
		}()
	}
	wg.Wait()

	if len(errors) == 0 {
		return ctx.JSON(http.StatusOK, echo.Map{
			"massage": fmt.Sprintf("All (%d) emails have been sent successfully.", len(students)),
		})
	} else {
		return ctx.JSON(http.StatusForbidden, echo.Map{
			"massage": fmt.Sprintf("%d emails have been sent successfully, %d failed.", len(students)-len(errors), len(errors)),
			"error":   errors[0],
		})
	}

}
