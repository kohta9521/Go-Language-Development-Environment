package main

import (
	"Go_Language_Development_Environment/todoapp-server/model"
	"Go_Language_Development_Environment/todoapp-server/router"

	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}