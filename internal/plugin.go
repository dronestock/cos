package internal

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/dronestock/cos/internal/internal/config"
	"github.com/dronestock/cos/internal/internal/step"
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type plugin struct {
	drone.Base

	Source config.Source `default:"SOURCE" json:"source,omitempty"`
	// 本身配置
	Cos config.Cos `default:"COS" json:"cos,omitempty"`
	// 密钥配置
	Secret config.Secret `default:"${SECRET}" json:"secret,omitempty"`

	cos *cos.Client
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Setup() (err error) {
	if client, coe := p.setupCos(); nil != coe {
		err = coe
	} else {
		p.cos = client
	}

	return
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewClear(&p.Source, p.cos)).Name("清理空间").Build(),
		drone.NewStep(step.NewUpload(&p.Source, &p.Cos, p.cos, p.Logger)).Name("上传文件").Build(),
		drone.NewStep(step.NewWebsite(&p.Cos, p.cos, p.Logger)).Name("静态网站").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("secret", p.Secret),
		field.New("cos", p.Cos),
	}
}

func (p *plugin) setupCos() (client *cos.Client, err error) {
	if endpoint, pe := url.Parse(p.Cos.Endpoint); nil != err {
		err = pe
	} else {
		transport := &cos.AuthorizationTransport{
			SecretID:  p.Secret.Id,
			SecretKey: p.Secret.Key,
			// nolint:gosec
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}
		client = cos.NewClient(&cos.BaseURL{BucketURL: endpoint}, &http.Client{
			Transport: transport,
		})
	}

	return
}
