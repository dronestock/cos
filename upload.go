package main

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func (p *plugin) upload() (undo bool, err error) {
	if paths, ae := gfx.All(p.Folder); nil != ae {
		err = ae
		p.Error("列举目录文件出错", field.New("folder", p.Folder), field.Error(err))
	} else {
		p.Debug("即将上传文件", field.New("paths", paths))
		for _, path := range paths {
			if err = p.uploadFile(path); nil != err {
				return
			}
		}
	}

	return
}

func (p *plugin) uploadFile(path string) (err error) {
	rel := ""
	pf := field.New("path", path)
	if _rel, re := filepath.Rel(p.Folder, path); nil != re {
		err = re
		p.Error("获取文件相对路径出错", pf, field.Error(err))
	} else {
		rel = _rel
	}
	if nil != err {
		return
	}

	paths := strings.Split(rel, string(filepath.Separator))
	if "" != p.Prefix {
		paths = append([]string{p.Prefix}, paths...)
	}
	if "" != p.Suffix {
		paths = append(paths, p.Suffix)
	}

	rel = strings.Join(paths, p.Separator)
	options := new(cos.MultiUploadOptions)
	options.CheckPoint = true
	if _, rsp, ue := p.cos.Object.MultiUpload(context.Background(), rel, path, options); nil != ue {
		err = ue
		p.Error("上传文件出错", pf, field.Error(err))
	} else if 200 != rsp.StatusCode {
		p.Warn("上传文件失败", pf)
	} else {
		p.Debug("文件上传成功", pf)
	}

	return
}
