package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/masutech16/pieceline/model"
	"github.com/masutech16/pieceline/router"
)

func main() {
	model.SetUp()
	e := echo.New()

	// session
	// よくわかっていないので調べる
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Authorization
	e.POST("/login", router.PostLogin)

	// groupを使ってmiddlewareを置く
	g := e.Group("/api/1.0")
	g.Use(router.AuthorityCheck)

	// twitter
	g.GET("/twitter/home", router.GetHomeTimeline)
	g.POST("/twitter/tweet", router.Tweet)
	g.PUT("/twitter/retweet", router.Retweet)
	g.PUT("/twitter/fav", router.FavTweet)
	g.POST("/twitter/reply", router.Reply)

	e.Logger.Fatal(e.Start(":1323"))
}
