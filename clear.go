package main

import (
	"strconv"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func clear(conf *config, _ *cos.Client) (err error) {
	var enable bool
	if enable, err = strconv.ParseBool(conf.Clear); nil != err || !enable {
		return
	}

	// 清理存储桶
	// client.Bucket.Get()

	return
}
