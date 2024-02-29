package controllers

import (
	"database/sql"
)

type Controller struct {
	DB *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		DB: db,
	}
}
