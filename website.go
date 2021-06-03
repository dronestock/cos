package main

import (
	"context"
	"strconv"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func website(conf *config, client *cos.Client) (err error) {
	var enable bool
	if enable, err = strconv.ParseBool(conf.Website.Enable); nil != err || !enable {
		return
	}

	var rsp *cos.BucketGetWebsiteResult
	if rsp, _, err = client.Bucket.GetWebsite(context.Background()); nil != err {
		return
	}

	if nil == rsp {
		_, err = client.Bucket.PutWebsite(context.Background(), &cos.BucketPutWebsiteOptions{
			Index: conf.Website.Index,
			Error: &cos.ErrorDocument{
				Key: conf.Website.Error,
			},
		})
	}

	return
}
