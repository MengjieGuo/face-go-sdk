package face

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/comeonjy/face-go-sdk/lib"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Face struct {
	Token *lib.AccessToken
}

func New(appKey string, appSecret string) *Face {
	return &Face{
		Token: lib.NewToken(appKey, appSecret),
	}
}

func (f *Face) tool(url string, v url.Values, res *Reply) (err error) {
	if err = f.Token.SetAccessToken(); err != nil {
		return
	}
	v.Add("access_token", f.Token.AccessToken)
	resp, err := http.PostForm(url, v)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bytes, res); err != nil {
		return
	}
	if res.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("%d:%s", res.ErrorCode, res.ErrorMsg))
	}
	return
}