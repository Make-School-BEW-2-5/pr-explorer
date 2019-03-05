package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Make-School-BEW-2-5/pr-explorer/database"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", GetData)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getAndSave(c echo.Context) error {
	database.SaveData(GetData)
	return c.String(http.StatusOK, "Exited")
}
