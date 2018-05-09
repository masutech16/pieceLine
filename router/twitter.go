package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/masutech16/pieceLine/model"
)

// GetHomeTimeline GET /home のハンドラ
func GetHomeTimeline(c echo.Context) error {
	tc := model.GetTwitterClient()
	tweets, err := tc.GetHomeTimeline()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tweets)
}
