package handlers

import "gorm.io/gorm"

type withDBHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *withDBHandler {
	h := withDBHandler{db}
	return &h
}

var WithDBHandler = &withDBHandler{}
