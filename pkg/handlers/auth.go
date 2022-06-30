package handlers

import (
	"net/http"
	"time"

	"github.com/Hesamsrk/golang-mail-server/pkg/auth"
	"github.com/Hesamsrk/golang-mail-server/pkg/config"
	"github.com/Hesamsrk/golang-mail-server/pkg/models"
	"github.com/Hesamsrk/golang-mail-server/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h withDBHandler) Login(ctx echo.Context) error {

	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	if err := ctx.Validate(user); err != nil {
		return err
	}

	var found models.User
	h.DB.Model(&user).Where("username = ?", user.Username).First(&found)

	if user.Username != found.Username ||  !utils.CheckPasswordHash(user.Password,found.Password)  {
		return echo.ErrUnauthorized
	}
	
	// Set custom claims
	claims := &auth.JwtCustomClaims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.LocalConfig.JWT_SECRET))

	h.DB.Model(&user).Where("username = ?", user.Username).Update("token", t)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h withDBHandler) Signup(ctx echo.Context) error {
	var user = &models.User{}
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	if err := ctx.Validate(user); err != nil {
		return err
	}

	var found models.User
	h.DB.Model(&user).Where("username = ?", user.Username).First(&found)

	if user.Username == found.Username {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"massage": "User already exists",
		})
	}

	updated := *user

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		return echo.ErrInternalServerError
	}
	updated.Password = password

	h.DB.Create(&updated)

	return ctx.JSON(http.StatusOK, user)
}
