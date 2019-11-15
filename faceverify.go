package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/comeonjy/face-go-sdk/lib"
	"io/ioutil"
	"net/http"
)

type VerifyRes struct {
	BaseReply
	Result struct {
		FaceLiveness float64 `json:"face_liveness"`
		FaceList     [] struct {
			FaceToken       string  `json:"face_token"`
			FaceProbability float64 `json:"face_probability"`
			Age             int64   `json:"age"`
			Beauty          float64 `json:"beauty"`
			Location        struct {
				Left     float64 `json:"left"`
				Top      float64 `json:"top"`
				Width    float64 `json:"width"`
				Height   float64 `json:"height"`
				Rotation float64 `json:"rotation"`
			} `json:"location"`
			Quality struct {
				Occlusion struct {
					LeftEye     float64 `json:"left_eye"`
					RightEye    float64 `json:"right_eye"`
					Nose        float64 `json:"nose"`
					Mouth       float64 `json:"mouth"`
					LeftCheek   float64 `json:"left_cheek"`
					RightCheek  float64 `json:"right_cheek"`
					ChinContour float64 `json:"chin_contour"`
				} `json:"occlusion"`
				Blur         float64 `json:"blur"`
				Illumination float64 `json:"illumination"`
				Completeness float64 `json:"completeness"`
			} `json:"quality"`
			FaceType struct {
				Type        string  `json:"type"`
				Probability float64 `json:"probability"`
			} `json:"face_type"`
		} `json:"face_list"`
	} `json:"result"`
}

// 活体检测
func (f *Face) FaceVerify(image, image_type, face_field, option string) (res *VerifyRes, err error) {
	res = &VerifyRes{}

	if err = f.Token.SetAccessToken(); err != nil {
		return
	}

	data := make(map[string]interface{})
	data["access_token"] = f.Token.AccessToken
	data["image"] = image
	data["image_type"] = image_type
	data["face_field"] = face_field
	data["option"] = option

	marshal, err := json.Marshal(data)
	req, err := http.NewRequest("POST", lib.FACE_VERIFY+"?access_token="+f.Token.AccessToken, bytes.NewReader([]byte("["+string(marshal)+"]")))
	client := &http.Client{}
	resp, err := client.Do(req)
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
		return nil, errors.New(fmt.Sprintf("%d:%s", res.ErrorCode, res.ErrorMsg))
	}

	return
}
