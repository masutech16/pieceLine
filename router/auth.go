package router

import (
	"net/http"

	"github.com/masutech16/pieceline/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// PostLogin パスワードとアカウント名で認証を行う
/*
 * @param {name: ユーザーネーム, password: パスワード}
 * 204: 成功。クッキーを渡す
 * 401: 認証失敗。
 *
 */
func PostLogin(c echo.Context) error {
	req := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "json format is wrong")
	}

	// model側で認証
	ok, err := model.Authorization(req.Name, req.Password)
	if err != nil {
		c.Logger().Errorf("Failed to Authorize: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to login")
	}
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password or Username is wrong")
	}

	// cookieを付ける
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 14,
		HttpOnly: true,
	}
	sess.Values["name"] = req.Name
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusNoContent)
}

// AuthorityCheck sessionを持っているかを確認する
func AuthorityCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := session.Get("session", c)
		if err != nil {
			c.Logger().Errorf("an error occurred while　checking session: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get session")
		}
		return next(c)
	}
}
