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

	//Validator:
	e.Validator = customValidator.New(validator.New())

	// middlewares: (log - preventPanic - CORS)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowCredentials: true,
	}))

	// docs:
	e.Static("/docs", "docs")
	//common:
	e.GET("/ping", handlers.Ping)
	e.PUT("/auth/signup", handlers.WithDBHandler.Signup)
	e.POST("/auth/login", handlers.WithDBHandler.Login)

	//Restricted paths:
	restricted := e.Group("/profile")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(config.LocalConfig.JWT_SECRET),
	}))
	//users:
	restricted.GET("/user/list", handlers.WithDBHandler.GetUserList)
	//students:
	restricted.PUT("/student/save", handlers.WithDBHandler.SaveStudent)
	restricted.DELETE("/student/remove", handlers.WithDBHandler.RemoveStudent)
	restricted.GET("/student/list", handlers.WithDBHandler.GetStudntsList)
	//classes:
	restricted.PUT("/class/save", handlers.WithDBHandler.SaveClass)
	restricted.DELETE("/class/remove", handlers.WithDBHandler.RemoveClass)
	restricted.GET("/class/list", handlers.WithDBHandler.GetClassesList)
	restricted.PUT("/email/SendClassDataToStudents", handlers.WithDBHandler.SendClassDataToStudents)
	e.Logger.Fatal(e.Start(":" + *PORT))
}
