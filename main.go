package main

import (
	"biblioteca/database"
	"biblioteca/routes"
	"biblioteca/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// database.InitConnection()
	// database.InitConnectionLite()
	database.InitConnectionPostgreSQL()
	utils.MigrateSchemas()
	routes.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
