package main

import (
	`context`
	`io/fs`
	`path/filepath`
	`strings`

	`github.com/tencentyun/cos-go-sdk-v5`
)

func upload(conf *config, client *cos.Client) (err error) {
	return filepath.WalkDir(conf.Folder, func(path string, dir fs.DirEntry, _ error) (err error) {
		if nil == dir || dir.IsDir() {
			return
		}

		var key string
		if key, err = filepath.Rel(conf.Folder, path); nil != err {
			return
		}
		paths := strings.Split(key, string(filepath.Separator))
		if "" != conf.Base {
			paths = append([]string{conf.Base}, paths...)
		}
		key = strings.Join(paths, conf.Separator)

		_, _, err = client.Object.MultiUpload(context.Background(), key, path, &cos.MultiUploadOptions{
			CheckPoint:         true,
			EnableVerification: true,
		})

		return
	})
}
