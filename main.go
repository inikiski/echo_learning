package main

import (
	"echo_learning/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,echo with gorm")
	})
	router.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
