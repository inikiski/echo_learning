package router

import (
	"echo_learning/api"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.GET("/save", api.SaveUser)
	e.GET("/get", api.GetUser)
	e.GET("/getAll", api.GetAll)
	e.GET("/update", api.UpdateUser)
	e.GET("/delete", api.DeleteUser)
}
