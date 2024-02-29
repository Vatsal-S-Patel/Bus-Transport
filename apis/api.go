package apis

import (
	"busproject/controllers"
	"database/sql"
)

type App struct {
	controller *controllers.Controller
}

func NewApp(db *sql.DB) *App {
	return &App{
		controller: controllers.NewController(db),
	}
}
