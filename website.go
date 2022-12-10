package main

import (
	"context"
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func (p *plugin) website() (undo bool, err error) {
	if undo = !p.Website.Enabled; undo {
		return
	}

	fields := gox.Fields[any]{
		field.New("website.index", p.Website.Index),
		field.New("website.error", p.Website.Error),
	}
	if result, rsp, ge := p.cos.Bucket.GetWebsite(context.Background()); nil != ge {
		err = ge
		p.Error("获取静态网站配置出错", fields.Connect(field.Error(err))...)
	} else if http.StatusOK != rsp.StatusCode {
		p.Warn("获取静态网站配置失败", fields...)
	} else {
		p.Debug("获取静态网站配置成功", fields...)
		err = p.putWebsite(result, rsp, fields)
	}

	return
}

func (p *plugin) putWebsite(result *cos.BucketGetWebsiteResult, rsp *cos.Response, fields gox.Fields[any]) (err error) {
	canPut := false
	if http.StatusNotFound != rsp.StatusCode {
		canPut = p.canPut(result, fields)
	} else {
		canPut = true
		p.Info("未发现静态网站配置", fields...)
	}
	if !canPut {
		return
	}

	options := new(cos.BucketPutWebsiteOptions)
	options.Index = p.Website.Index
	options.Error = &cos.ErrorDocument{
		Key: p.Website.Error,
	}

	if rsp, pe := p.cos.Bucket.PutWebsite(context.Background(), options); nil != pe {
		err = pe
		p.Error("配置静态网站配置出错", fields.Connect(field.Error(err))...)
	} else if 200 != rsp.StatusCode {
		p.Warn("配置静态网站配置失败", fields...)
	} else {
		p.Debug("配置静态网站配置成功", fields...)
	}

	return
}

func (p *plugin) canPut(result *cos.BucketGetWebsiteResult, fields gox.Fields[any]) (canPut bool) {
	if canPut = p.Website.Index != result.Index || p.Website.Error != result.Error.Key; canPut {
		p.Debug("静态网站配置和现在的配置不同，需要更新", fields...)
	} else {
		p.Debug("静态网站配置和现在的配置相同，不用更新", fields...)
	}

	return
}
