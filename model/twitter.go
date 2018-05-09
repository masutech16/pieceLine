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
	}

	twitterClient struct {
		api *anaconda.TwitterApi
	}
)

func (tc *twitterClient) PostTweet(status string) error {
	_, err := tc.api.PostTweet(status, nil)
	return err
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
