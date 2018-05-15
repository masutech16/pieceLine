package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masutech16/pieceLine/model"
	"github.com/masutech16/pieceLine/router"
)

func main() {
	model.SetUp()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/home", router.GetHomeTimeline)
	e.Logger.Fatal(e.Start(":1323"))
}
