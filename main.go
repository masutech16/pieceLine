package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/masutech16/pieceline/model"
	"github.com/masutech16/pieceline/router"
)

func main() {
	model.SetUp()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// twitter
	e.GET("/twitter/home", router.GetHomeTimeline)
	e.POST("/twitter/tweet", router.Tweet)
	e.PUT("/twitter/retweet", router.Retweet)
	e.PUT("/twitter/fav", router.FavTweet)
	e.POST("/twitter/reply", router.Reply)
	e.Logger.Fatal(e.Start(":1323"))
}
