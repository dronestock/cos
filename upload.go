package main

import (
	`context`
	`path/filepath`
	`strings`

	`github.com/storezhang/gfx`
	`github.com/tencentyun/cos-go-sdk-v5`
)

func (p *plugin) upload() (undo bool, err error) {
	var paths []string
	if paths, err = gfx.All(p.Folder); nil != err {
		return
	}

	for _, path := range paths {
		if err = p.uploadFile(path); nil != err {
			return
		}
	}

	return
}

func (p *plugin) uploadFile(path string) (err error) {
	var rel string
	if rel, err = filepath.Rel(p.Folder, path); nil != err {
		return
	}

	paths := strings.Split(rel, string(filepath.Separator))
	if `` != p.Prefix {
		paths = append([]string{p.Prefix}, paths...)
	}
	if `` != p.Suffix {
		paths = append(paths, p.Suffix)
	}

	rel = strings.Join(paths, p.Separator)
	_, _, err = p.cos.Object.MultiUpload(context.Background(), rel, path, &cos.MultiUploadOptions{
		CheckPoint: true,
	})

	return
}
