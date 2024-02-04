package main

import (
	"biblioteca/database"
	"biblioteca/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	database.InitConnectionPostgreSQL()
	// utils.MigrateSchemas()
	routes.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
