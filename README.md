# face-go-sdk
百度人脸识别SDK

- [x] 人脸库管理  
- [x] 在线活体检测 
- [ ] 人脸搜索 
- [ ] 人脸对比 
- [ ] 人脸检测 
- [ ] 身份验证 

##示例
```go
var f = New(APP_KEY, APP_SECRET)
if res, err := f.GetUser("1", "demo1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
```

