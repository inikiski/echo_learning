package api

import (
	"echo_learning/dao"
	"github.com/labstack/echo/v4"
	"time"
)

func SaveUser(e echo.Context) error {
	user := &dao.User{
		Username:   "gogogo",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.SaveUser(user)
	return e.JSON(200, user)
}

func GetUser(e echo.Context) error {
	user := dao.GetById(1)
	return e.JSON(200, user)
}

func GetAll(e echo.Context) error {
	user := dao.GetAll()
	return e.JSON(200, user)
}

func UpdateUser(e echo.Context) error {
	dao.UpdateUser(1)
	user := dao.GetById(1)
	return e.JSON(200, user)
}

func DeleteUser(e echo.Context) error {
	dao.DeleteUser(1)
	user := dao.GetById(1)
	return e.JSON(200, user)
}
