package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

const (
	twitterAPIRoot = "https://api.twitter.com"
)

type (
	twitterClient struct {
		api *anaconda.TwitterApi
	}

	// Tweet Tweetの構造体
	Tweet struct {
		Text          string `json:"text"`
		CreatedAt     string `json:"created_at"`
		FavoriteCount int    `json:"favorite_count"`
		Favorited     bool   `json:"favorited"`
		RetweetCount  int    `json:"retweet_count"`
		Retweeted     bool   `json:"retweeted"`
		UserID        string `json:"user_id"`
		DisplayName   string `json:"display_name"`
		IconURL       string `json:"icon_url"`
	}
)

// PostTweet ツイートを投稿する
func PostTweet(status string) (*Tweet, error) {
	tw, err := client.api.PostTweet(status, nil)
	if err != nil {
		return nil, err
	}
	return formatTweet(tw), nil
}

// GetHomeTimeline Homeディレクトリのツイート最新20件を取得する
func GetHomeTimeline() ([]*Tweet, error) {
	tweets, err := client.api.GetHomeTimeline(nil)
	if err != nil {
		return nil, err
	}
	var val []*Tweet
	for _, v := range tweets {
		val = append(val, formatTweet(v))
	}
	return val, nil
}

// FavTweet ツイートをふぁぼる
func FavTweet(id int64) (*Tweet, error) {
	tw, err := client.api.Favorite(id)
	if err != nil {
		return nil, err
	}

	return formatTweet(tw), nil
}

// Retweet リツイートする
func Retweet(id int64) (*Tweet, error) {

	//trim_userはとりあえず常にtrueで
	tw, err := client.api.Retweet(id, true)
	if err != nil {
		return nil, err
	}

	return formatTweet(tw), nil
}

// Reply 指定したツイートにリプライを送る
func Reply(status string, id string) (*Tweet, error) {
	v := url.Values{}
	v.Add("in_reply_to_status_id", id)
	tw, err := client.api.PostTweet(status, v)
	if err != nil {
		return nil, err
	}

	return formatTweet(tw), nil
}

// PostRequestToken /oauth/request_tokenにポストする
func PostRequestToken() error {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("POST", twitterAPIRoot+"/oauth/request_token", nil)
	if err != nil {
		return fmt.Errorf("Failed to make request")
	}

	// TODO: Authorization Headerの作成

	req.Header.Add("Authorization", url.QueryEscape("http://127.0.0.1:1323"))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("response returned an error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
	return nil
}

func formatTweet(raw anaconda.Tweet) *Tweet {
	return &Tweet{
		Text:          raw.FullText,
		CreatedAt:     raw.CreatedAt,
		FavoriteCount: raw.FavoriteCount,
		Favorited:     raw.Favorited,
		RetweetCount:  raw.RetweetCount,
		Retweeted:     raw.Retweeted,
		UserID:        raw.User.ScreenName,
		DisplayName:   raw.User.Name,
		IconURL:       raw.User.ProfileImageURL,
	}

}
