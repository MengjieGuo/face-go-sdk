package face

/**
 * API文档 https://cloud.baidu.com/doc/FACE/s/Xk25rdcy6
 */

import (
	"encoding/json"
	"github.com/comeonjy/face-go-sdk/lib"
)

type FaceMatch struct {
	Image           string
	ImageType       string
	FaceType        string
	QualityControl  string
	LivenessControl string
}

type MatchRes struct {
	Score    float64 `json:"score"`
	FaceList []struct {
		FaceToken string `json:"face_token"`
	} `json:"face_list"`
}

// 人脸对比
func (f *Face) MatchFace(search []*FaceMatch) (res *MatchRes, err error) {
	res = &MatchRes{}
	var data []map[string]string
	for _, v := range search {
		temp := map[string]string{
			"image":      v.Image,
			"image_type": v.ImageType,
		}
		if v.FaceType != "" {
			temp["face_type"] = v.FaceType
		}
		if v.QualityControl != "" {
			temp["quality_control"] = v.QualityControl
		}
		if v.LivenessControl != "" {
			temp["liveness_control"] = v.LivenessControl
		}
		data = append(data, temp)
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = f.PostJson(lib.FACE_MATCH, marshal, &Reply{Result: res})

	return
}
