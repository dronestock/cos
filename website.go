package main

import (
	`context`
	`net/http`

	`github.com/tencentyun/cos-go-sdk-v5`
)

func (p *plugin) website() (undo bool, err error) {
	if undo = p.websiteEnabled(); undo {
		return
	}

	if _, _, err = p.cos.Bucket.GetWebsite(context.Background()); nil != err {
		if http.StatusNotFound == err.(*cos.ErrorResponse).Response.StatusCode {
			_, err = p.cos.Bucket.PutWebsite(context.Background(), &cos.BucketPutWebsiteOptions{
				Index: p.WebsiteIndex,
				Error: &cos.ErrorDocument{
					Key: p.WebsiteError,
				},
			})
		}
	}

	return
}
