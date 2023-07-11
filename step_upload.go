package main

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type stepUpload struct {
	*plugin

	paths []string
}

func newUploadStep(plugin *plugin) *stepUpload {
	return &stepUpload{
		plugin: plugin,
	}
}

func (u *stepUpload) Runnable() (runnable bool) {
	if paths, ae := gfx.All(u.Folder); nil == ae || 0 != len(paths) {
		runnable = true
		u.paths = paths
	}

	return
}

func (u *stepUpload) Run(_ context.Context) (err error) {
	for _, path := range u.paths {
		if err = u.uploadFile(path); nil != err {
			return
		}
	}

	return
}

func (u *stepUpload) uploadFile(path string) (err error) {
	rel := ""
	pf := field.New("path", path)
	if _rel, re := filepath.Rel(u.Folder, path); nil != re {
		err = re
		u.Error("获取文件相对路径出错", pf, field.Error(err))
	} else {
		rel = _rel
	}
	if nil != err {
		return
	}

	paths := strings.Split(rel, string(filepath.Separator))
	if "" != u.Prefix {
		paths = append([]string{u.Prefix}, paths...)
	}
	if "" != u.Suffix {
		paths = append(paths, u.Suffix)
	}

	rel = strings.Join(paths, u.Separator)
	options := new(cos.MultiUploadOptions)
	options.CheckPoint = true
	if _, rsp, ue := u.cos.Object.MultiUpload(context.Background(), rel, path, options); nil != ue {
		err = ue
		u.Error("上传文件出错", pf, field.Error(err))
	} else if 200 != rsp.StatusCode {
		u.Warn("上传文件失败", pf)
	} else {
		u.Debug("文件上传成功", pf)
	}

	return
}
