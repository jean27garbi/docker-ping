package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jean27garbi/docker-ping/cmd/db"
	"github.com/jean27garbi/docker-ping/cmd/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Server
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init db

	data, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialize the store: %s", err)
	}
	// Routers

	e.GET("/", func(c echo.Context) error {
		return handlers.RootHandler(data, c)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.POST("/send", func(c echo.Context) error {
		return handlers.SendHandler(data, c)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}
