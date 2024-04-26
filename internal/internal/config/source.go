package config

type Source struct {
	// 本地上传目录
	Folder string `default:"${FOLDER=.}" json:"folder,omitempty"`
}
