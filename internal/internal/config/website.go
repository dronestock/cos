package config

type Website struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled"`
	// 主页
	Index string `default:"index.html" json:"index"`
	// 错误页
	Error string `json:"error"`
}
