package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Hesamsrk/golang-mail-server/pkg/auth"
	"github.com/Hesamsrk/golang-mail-server/pkg/config"
	"github.com/Hesamsrk/golang-mail-server/pkg/customValidator"
	"github.com/Hesamsrk/golang-mail-server/pkg/db"
	"github.com/Hesamsrk/golang-mail-server/pkg/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	config.LocalConfig = &c
	DB := db.Init()
	handlers.WithDBHandler = handlers.New(DB)
}

func main() {

	PORT := flag.String("port", "9090", "The HTTP port to listen on")
	flag.Parse()
	fmt.Printf("configs: %#v", *config.LocalConfig)

	e := echo.New()
	e.Validator = customValidator.New(validator.New())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	e.GET("/", handlers.Landing)
	e.GET("/ping", handlers.Ping)
	e.PUT("/auth/signup", handlers.WithDBHandler.Signup)
	e.POST("/auth/login", handlers.WithDBHandler.Login)

	restricted := e.Group("/profile")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(config.LocalConfig.JWT_SECRET),
	}))
	restricted.GET("/user/list", handlers.WithDBHandler.GetUserList)
	e.Any("*", handlers.NotFound)
	e.Logger.Fatal(e.Start(":" + *PORT))
}
