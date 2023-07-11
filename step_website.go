package main

import (
	"context"
	"net/http"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type stepWebsite struct {
	*plugin
}

func newWebsiteStep(plugin *plugin) *stepWebsite {
	return &stepWebsite{
		plugin: plugin,
	}
}

func (w *stepWebsite) Runnable() bool {
	return w.Website.Enabled
}

func (w *stepWebsite) Run(_ context.Context) (err error) {
	fields := gox.Fields[any]{
		field.New("website.index", w.Website.Index),
		field.New("website.error", w.Website.Error),
	}
	if result, rsp, ge := w.cos.Bucket.GetWebsite(context.Background()); nil != ge {
		err = ge
		w.Error("获取静态网站配置出错", fields.Add(field.Error(err))...)
	} else if http.StatusOK != rsp.StatusCode {
		w.Warn("获取静态网站配置失败", fields...)
	} else {
		w.Debug("获取静态网站配置成功", fields...)
		err = w.put(result, rsp, fields)
	}

	return
}

func (w *stepWebsite) put(result *cos.BucketGetWebsiteResult, rsp *cos.Response, fields gox.Fields[any]) (err error) {
	canPut := false
	if http.StatusNotFound != rsp.StatusCode {
		canPut = w.can(result, fields)
	} else {
		canPut = true
		w.Info("未发现静态网站配置", fields...)
	}
	if !canPut {
		return
	}

	options := new(cos.BucketPutWebsiteOptions)
	options.Index = w.Website.Index
	options.Error = &cos.ErrorDocument{
		Key: w.Website.Error,
	}

	if rsp, pe := w.cos.Bucket.PutWebsite(context.Background(), options); nil != pe {
		err = pe
		w.Error("配置静态网站配置出错", fields.Add(field.Error(err))...)
	} else if 200 != rsp.StatusCode {
		w.Warn("配置静态网站配置失败", fields...)
	} else {
		w.Debug("配置静态网站配置成功", fields...)
	}

	return
}

func (w *stepWebsite) can(result *cos.BucketGetWebsiteResult, fields gox.Fields[any]) (canPut bool) {
	if canPut = w.Website.Index != result.Index || w.Website.Error != result.Error.Key; canPut {
		w.Debug("静态网站配置和现在的配置不同，需要更新", fields...)
	} else {
		w.Debug("静态网站配置和现在的配置相同，不用更新", fields...)
	}

	return
}
