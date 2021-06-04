package main

import (
	`crypto/tls`
	`encoding/json`
	`fmt`
	`net/http`
	`net/url`

	`github.com/mcuadros/go-defaults`
	`github.com/tencentyun/cos-go-sdk-v5`
)

func main() {
	var err error
	// 取各种参数
	conf := &config{}
	conf.SecretId = env("SECRET_ID")
	conf.SecretKey = env("SECRET_KEY")
	conf.Folder = env("FOLDER")
	conf.BaseUrl = env("BASE_URL")
	conf.Clear = env("CLEAR")
	conf.Clear = env("BASE")
	conf.Website.Enable = env("WEBSITE_ENABLE")
	conf.Website.Index = env("WEBSITE_INDEX")
	conf.Website.Error = env("WEBSITE_ERROR")
	defaults.SetDefaults(conf)
	fmt.Println(json.Marshal(conf))

	var bucketUrl *url.URL
	if bucketUrl, err = url.Parse(conf.BaseUrl); nil != err {
		panic(err)
	}

	client := cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretId,
			SecretKey: conf.SecretKey,
			// nolint:gosec
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	})

	// 清空存储桶
	if err = clear(conf, client); nil != err {
		panic(err)
	}

	// 上传文件
	if err = upload(conf, client); nil != err {
		panic(err)
	}

	// 静态网站
	if err = website(conf, client); nil != err {
		panic(err)
	}
}
