package main

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type plugin struct {
	drone.Base

	// 本地上传目录
	Folder string `default:"${FOLDER=.}"`

	// 授权，类型于用户名
	SecretId string `default:"${SECRET_ID}" validate:"required"`
	// 授权，类型于密码
	SecretKey string `default:"${SECRET_KEY}" validate:"required"`
	// 存储桶地址
	BaseUrl string `default:"${BASE_URL}" validate:"required,url"`

	// 分隔符
	Separator string `default:"${SEPARATOR=/}"`
	// 是否清空存储桶
	Clear bool `default:"${CLEAR=true}"`
	// 路径前缀，所有文件上传都会在这上面加上前缀
	Prefix string `default:"${PREFIX}"`
	// 路径后缀，所有文件上传都会在这上面加上后缀
	Suffix string `default:"${SUFFIX}"`

	// 静态网站
	Website websiteConfig `default:"${WEBSITE}"`

	cos *cos.Client
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (err error) {
	var bucketUrl *url.URL
	if bucketUrl, err = url.Parse(p.BaseUrl); nil != err {
		panic(err)
	}

	p.cos = cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  p.SecretId,
			SecretKey: p.SecretKey,
			// nolint:gosec
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	})

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(newClearStep(p)).Name("清理空间").Build(),
		drone.NewStep(newUploadStep(p)).Name("上传文件").Build(),
		drone.NewStep(newWebsiteStep(p)).Name("静态网站").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("folder", p.Folder),

		field.New("secret.id", p.SecretId),
		field.New("base.url", p.BaseUrl),

		field.New("separator", p.Separator),
		field.New("clear", p.Clear),
		field.New("prefix", p.Prefix),
		field.New("suffix", p.Suffix),

		field.New("websiteConfig.index", p.Website.Index),
		field.New("websiteConfig.error", p.Website.Error),
	}
}
