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
	// TwitterClient ツイッターと通信するインターフェース
	TwitterClient interface {
		PostTweet(string) error
		GetHomeTimeline() ([]*Tweet, error)
	}

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

func (tc *twitterClient) PostTweet(status string) error {
	_, err := tc.api.PostTweet(status, nil)
	return err
}

// GetHomeTimeline Homeディレクトリのツイート最新20件を取得する
func (tc *twitterClient) GetHomeTimeline() ([]*Tweet, error) {
	tweets, err := tc.api.GetHomeTimeline(nil)
	if err != nil {
		return nil, err
	}
	var val []*Tweet
	for _, v := range tweets {
		t := &Tweet{
			Text:          v.FullText,
			CreatedAt:     v.CreatedAt,
			FavoriteCount: v.FavoriteCount,
			Favorited:     v.Favorited,
			RetweetCount:  v.RetweetCount,
			Retweeted:     v.Retweeted,
			UserID:        v.User.ScreenName,
			DisplayName:   v.User.Name,
			IconURL:       v.User.ProfileImageURL,
		}
		val = append(val, t)
	}
	return val, nil
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
