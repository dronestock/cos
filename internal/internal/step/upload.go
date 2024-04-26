package step

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/dronestock/cos/internal/internal/config"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Upload struct {
	source *config.Source
	config *config.Cos
	paths  []string
	cos    *cos.Client
	logger log.Logger
}

func NewUpload(source *config.Source, config *config.Cos, cos *cos.Client, logger log.Logger) *Upload {
	return &Upload{
		source: source,
		config: config,
		cos:    cos,
		logger: logger,
	}
}

func (u *Upload) Runnable() (runnable bool) {
	if paths, ae := gfx.All(u.source.Folder); nil == ae || 0 != len(paths) {
		runnable = true
		u.paths = paths
	}

	return
}

func (u *Upload) Run(_ *context.Context) (err error) {
	for _, path := range u.paths {
		if err = u.uploadFile(path); nil != err {
			return
		}
	}

	return
}

func (u *Upload) uploadFile(path string) (err error) {
	rel := ""
	pf := field.New("path", path)
	if _rel, re := filepath.Rel(u.source.Folder, path); nil != re {
		err = re
		u.logger.Error("获取文件相对路径出错", pf, field.Error(err))
	} else {
		rel = _rel
	}
	if nil != err {
		return
	}

	paths := strings.Split(rel, string(filepath.Separator))
	if "" != u.config.Prefix {
		paths = append([]string{u.config.Prefix}, paths...)
	}
	if "" != u.config.Suffix {
		paths = append(paths, u.config.Suffix)
	}

	rel = strings.Join(paths, u.config.Separator)
	options := new(cos.MultiUploadOptions)
	options.CheckPoint = true
	if _, rsp, ue := u.cos.Object.MultiUpload(context.Background(), rel, path, options); nil != ue {
		err = ue
		u.logger.Error("上传文件出错", pf, field.Error(err))
	} else if 200 != rsp.StatusCode {
		u.logger.Warn("上传文件失败", pf)
	} else {
		u.logger.Debug("文件上传成功", pf)
	}

	return
}
