package main

type websiteConfig struct {
	// 是否开户
	Enabled bool `default:"true" json:"enabled"`
	// 主页
	Index string `default:"index.html" json:"index"`
	// 错误页
	Error string `default:"error.html" json:"error"`
}
