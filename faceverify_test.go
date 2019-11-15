package face

import (
	"fmt"
	"testing"
)

func TestFace_FaceVerify(t *testing.T) {
	if res,err:=f.FaceVerify("600d3bf0edb0742e90a8e20e2d7d28f5","FACE_TOKEN","age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type","COMMON");err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%+v \n",res)
	}
}