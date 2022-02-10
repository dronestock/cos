package main

import (
	`crypto/tls`
	`net/http`
	`net/url`
	`strings`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/tencentyun/cos-go-sdk-v5`
)

type plugin struct {
	drone.PluginBase

	// 本地上传目录
	Folder string `default:"${PLUGIN_Folder=${Folder=.}}"`

	// 授权，类型于用户名
	SecretId string `default:"${PLUGIN_SECRET_ID=${SECRET_ID}}" validate:"required"`
	// 授权，类型于密码
	SecretKey string `default:"${PLUGIN_SECRET_KEY=${SECRET_KEY}}" validate:"required"`
	// 存储桶地址
	BaseUrl string `default:"${PLUGIN_BASE_URL=${BASE_URL}}" validate:"required,url"`

	// 分隔符
	Separator string `default:"${PLUGIN_SEPARATOR=${SEPARATOR=/}}"`
	// 是否清空存储桶
	Clear bool `default:"${PLUGIN_CLEAR=${CLEAR=true}}"`
	// 路径前缀，所有文件上传都会在这上面加上前缀
	Prefix string `default:"${PLUGIN_PREFIX=${PREFIX}}"`
	// 路径后缀，所有文件上传都会在这上面加上后缀
	Suffix string `default:"${PLUGIN_SUFFIX=${SUFFIX}}"`

	// 静态网站主页
	WebsiteIndex string `default:"${PLUGIN_WEBSITE_INDEX=${WEBSITE_INDEX=index.html}}"`
	// 静态网站错误页
	WebsiteError string `default:"${PLUGIN_WEBSITE_ERROR=${WEBSITE_ERROR=error.html}}"`

	cos *cos.Client
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (unset bool, err error) {
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

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.clear, drone.Name(`清理存储空间`)),
		drone.NewStep(p.upload, drone.Name(`上传文件`)),
		drone.NewStep(p.website, drone.Name(`配置静态网站`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Folder),

		field.String(`secret.id`, p.SecretId),
		field.String(`base.url`, p.BaseUrl),

		field.String(`separator`, p.Separator),
		field.Bool(`clear`, p.Clear),
		field.String(`prefix`, p.Prefix),
		field.String(`suffix`, p.Suffix),

		field.String(`website.index`, p.WebsiteIndex),
		field.String(`website.error`, p.WebsiteError),
	}
}

func (p *plugin) websiteEnabled() bool {
	return `` == strings.TrimSpace(p.WebsiteIndex) && `` == strings.TrimSpace(p.WebsiteError)
}
