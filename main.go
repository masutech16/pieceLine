package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masutech16/pieceLine/model"
)

func main() {
	model.SetUp()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	tc := model.GetTwitterClient()
	err := tc.PostTweet("test")
	if err != nil {
		return fmt.Errorf("an error occurred in posttweet: %v", err)
	}
	return c.String(http.StatusOK, "Hello World!")
}
