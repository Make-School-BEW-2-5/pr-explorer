package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	bolt "go.etcd.io/bbolt"

	"./database"
)

var DB *bolt.DB

func main() {
	// Echo instance
	e := echo.New()

	DB = database.Init()
	defer DB.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", getAndSave)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getAndSave(c echo.Context) error {
	database.SaveData(DB, GetData())
	return c.String(http.StatusOK, "Exited")
}
