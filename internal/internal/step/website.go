package step

import (
	"context"
	"net/http"

	"github.com/dronestock/cos/internal/internal/config"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Website struct {
	config *config.Wrapper
	cos    *cos.Client
	logger log.Logger
}

func NewWebsite(config *config.Wrapper, cos *cos.Client, logger log.Logger) *Website {
	return &Website{
		config: config,
		cos:    cos,
		logger: logger,
	}
}

func (w *Website) Runnable() bool {
	return *w.config.Website.Enabled
}

func (w *Website) Run(_ *context.Context) (err error) {
	fields := gox.Fields[any]{
		field.New("website.index", w.config.Website.Index),
		field.New("website.error", w.config.Website.Error),
	}
	if result, rsp, ge := w.cos.Bucket.GetWebsite(context.Background()); nil != ge {
		err = ge
		w.logger.Error("获取静态网站配置出错", fields.Add(field.Error(err))...)
	} else if http.StatusOK != rsp.StatusCode {
		w.logger.Warn("获取静态网站配置失败", fields...)
	} else {
		w.logger.Debug("获取静态网站配置成功", fields...)
		err = w.put(result, rsp, fields)
	}

	return
}

func (w *Website) put(result *cos.BucketGetWebsiteResult, rsp *cos.Response, fields gox.Fields[any]) (err error) {
	canPut := false
	if http.StatusNotFound != rsp.StatusCode {
		canPut = w.can(result, fields)
	} else {
		canPut = true
		w.logger.Info("未发现静态网站配置", fields...)
	}
	if !canPut {
		return
	}

	options := new(cos.BucketPutWebsiteOptions)
	options.Index = w.config.Website.Index
	options.Error = &cos.ErrorDocument{
		Key: w.config.Website.Error,
	}
	if pwr, pe := w.cos.Bucket.PutWebsite(context.Background(), options); nil != pe {
		err = pe
		w.logger.Error("配置静态网站配置出错", fields.Add(field.Error(err))...)
	} else if 200 != pwr.StatusCode {
		w.logger.Warn("配置静态网站配置失败", fields...)
	} else {
		w.logger.Debug("配置静态网站配置成功", fields...)
	}

	return
}

func (w *Website) can(result *cos.BucketGetWebsiteResult, fields gox.Fields[any]) (can bool) {
	can = w.config.Website.Index != result.Index
	if !can && "" != w.config.Website.Error {
		can = nil != result.Error && w.config.Website.Error != result.Error.Key
	}
	if can {
		w.logger.Debug("静态网站配置和现在的配置不同，需要更新", fields...)
	} else {
		w.logger.Debug("静态网站配置和现在的配置相同，不用更新", fields...)
	}

	return
}
