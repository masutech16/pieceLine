package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/masutech16/pieceLine/model"
)

// GetHomeTimeline GET /home のハンドラ
/*
 * 200: 正常に取得
 * 401: 認証失敗
 */
func GetHomeTimeline(c echo.Context) error {
	tweets, err := model.GetHomeTimeline()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tweets)
}

// Tweet POST /twitter/tweet のハンドラ
/*
 * @param {status: "tweetしたい文字列"}
 * 204: tweet成功
 *		Tweetオブジェクトを返します
 * 401: 認証失敗
 */
func Tweet(c echo.Context) error {
	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON format is wrong")
	}

	if len(req.Status) > 140 {
		return echo.NewHTTPError(http.StatusBadRequest, "140 over!")
	}

	tw, err := model.PostTweet(req.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to post tweet")
	}

	return c.JSON(http.StatusOK, tw)
}

// Retweet PUT /twitter/retweetのハンドラ
/*
 * @param {id: retweetするtweetID}
 * 200: 成功。Retweetオブジェクトを返す
 * 401: 認証失敗
 *
 */
func Retweet(c echo.Context) error {
	var req struct {
		ID int64 `json:"id"`
	}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON format is wrong")
	}

	tw, err := model.Retweet(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retweet")
	}

	return c.JSON(http.StatusOK, tw)
}

// FavTweet PUT /twitter/favのハンドラ
/*
 * @param {id: favするtweetID}
 * 204: 成功。Retweetオブジェクトを返す
 * 401: 認証失敗
 *
 */
func FavTweet(c echo.Context) error {
	var req struct {
		ID string `json:"id"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON format is wrong")
	}

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return err
	}
	tw, err := model.FavTweet(int64(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to make favorite")
	}

	return c.JSON(http.StatusOK, tw)
}

// Reply POST /twitter/replyのハンドラ
/*
 * @param {id: replyするtweetID, status: ツイート内容}
 * 204: 成功。Retweetオブジェクトを返す
 * 401: 認証失敗
 *
 */
func Reply(c echo.Context) error {
	var req struct {
		ID     int64  `json:"id"`
		Status string `json:"status"`
	}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON format is wrong")
	}

	if len(req.Status) > 140 {
		return echo.NewHTTPError(http.StatusBadRequest, "140 over!")
	}

	tw, err := model.PostTweet(req.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to post tweet")
	}

	return c.JSON(http.StatusOK, tw)
}
