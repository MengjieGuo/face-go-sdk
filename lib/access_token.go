package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type AccessToken struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	UpdateTimeStamp  int64  `json:"update_time_stamp"`
}

var appKey, appSecret string

func NewToken(app_key string, app_secret string) *AccessToken {
	appKey = app_key
	appSecret = app_secret
	return &AccessToken{}
}

func (t *AccessToken) SetAccessToken() (err error) {

	if time.Now().Unix()-t.UpdateTimeStamp < t.ExpiresIn {
		return nil
	}
	resp, err := http.PostForm(ACCESS_TOKEN_URL, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {appKey},
		"client_secret": {appSecret},
	})
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(bytes, t); err != nil {
		return
	}
	if t.Error != "" {
		return errors.New(fmt.Sprintf("%s:%s", t.Error, t.ErrorDescription))
	}
	t.UpdateTimeStamp = time.Now().Unix()

	return nil
}
