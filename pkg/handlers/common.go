package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Landing(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "Mail server: request /ping")
}

func Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}


var NotFound echo.HandlerFunc = func(ctx echo.Context) error {
	return ctx.String(http.StatusNotFound, "Not found!")
}
