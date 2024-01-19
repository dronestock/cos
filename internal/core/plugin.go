package core

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/dronestock/cos/internal/config"
	"github.com/dronestock/cos/internal/step"
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Plugin struct {
	drone.Base
	config.Wrapper

	refresh config.Refresh `default:"${REFRESH}"`

	cos *cos.Client
	cdn *cdn.Client
}

func NewPlugin() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Setup() (err error) {
	if endpoint, pe := url.Parse(p.Endpoint); nil != err {
		err = pe
	} else if cdnClient, cde := p.setupCdn(); nil != cde {
		p.cdn = cdnClient
	} else {
		p.cos = cos.NewClient(&cos.BaseURL{BucketURL: endpoint}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  p.Secret.Id,
				SecretKey: p.Secret.Key,
				// nolint:gosec
				Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			},
		})
	}

	return
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewClear(&p.Wrapper, p.cos)).Name("清理空间").Build(),
		drone.NewStep(step.NewUpload(&p.Wrapper, p.cos, p.Logger)).Name("上传文件").Build(),
		drone.NewStep(step.NewWebsite(&p.Wrapper, p.cos, p.Logger)).Name("静态网站").Build(),
		drone.NewStep(step.NewRefresh(&p.refresh, p.cdn)).Name("刷新预热").Build(),
	}
}

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("folder", p.Folder),
		field.New("secret", p.Secret),
		field.New("endpoint", p.Endpoint),
		field.New("separator", p.Separator),
		field.New("clear", p.Clear),
		field.New("prefix", p.Prefix),
		field.New("suffix", p.Suffix),
		field.New("website", p.Website),
	}
}

func (p *Plugin) setupCdn() (client *cdn.Client, err error) {
	credential := common.NewCredential(p.Secret.Id, p.Secret.Key)
	_profile := profile.NewClientProfile()

	return cdn.NewClient(credential, p.refresh.Regin, _profile)
}
