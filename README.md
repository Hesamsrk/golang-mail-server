# Golang-mail-server

This project is a golang server implemented using [Echo](https://echo.labstack.com/) which can help school or university managers to send students their scores by email. 

- This server uses `JWT` authentication.
- This project comes with a `PostgreSQL` DBMS which you can run and use it following the instructions below.
- I also used [GORM](https://gorm.io/), [ZOHO mail client](https://www.zoho.com/), [smtp](https://pkg.go.dev/net/smtp) and [go-playground-validator](github.com/go-playground/validator/v1)

# How to run it

- To make running CLI commands easier, I have used `makefile`. So if you don't have `makefile` installed on your system, please go to `./makefile` and run these commands manually by copy pasting.
- I also used docker for running `Postgres` DBMS. So you also need docker installed on your system.
- Paying attention to last points, run:

```shell

make prepare
# Alternatively:
# go mod tidy

make db-run
# Alternatively: (Add `sudo` on unix based systems)
# docker run --name golang-email-server \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=staging \
	-e POSTGRES_USER=pg \
	-p 9099:5432 \
	-d postgres

make dev
# Alternatively:
# 	 go run main.go --port 9090

```
- Now DBMS is running on https://localhost:9099 and server is running on https://localhost:9090.
- Access docs: https://localhost:9090/docs

# Documentations

## API

You can access API documentation by accessing http://localhost:9090/docs after running the program on that specific port. You can also run the APIs on [Insomnia](https://insomnia.rest/), using `Run in insomnia` button, if you have it  already installed on your machine.

## Schemas

```golang

// Base model:
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User:
type User struct {
	Model
	Username string `validate:"required,min=5" gorm:"unique"`
	Password string `validate:"required,min=5"`
	Token    string `validate:"omitempty"`
}

// Student:
type Student struct {
	Model
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `gorm:"unique" validate:"required,email"`
	Score     uint8  `validate:"required,gte=0,lte=20"`
	ClassID   uint
	Class     Class `gorm:"references:ID"`
}

// Class:

type Class struct {
	Model
	Field   string
	Teacher string
}
```
> ## Bonus features
> 1. Using docker for running database. ✅
> 2. Using `go routines` and `syng.waitGroup` for sending mails concurrently. ✅