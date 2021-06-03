package main

type config struct {
	// 授权，类型于用户名
	SecretId string
	// 授权，类型于密码
	SecretKey string
	// 存储桶地址
	BaseUrl string
	// 本地上传目录
	Folder string `default:"."`
	// 分隔符
	Separator string `default:"/"`
	// 是否清空存储桶
	Clear string `default:"false"`
	// 静态网站
	Website struct {
		// 是否启用
		Enable string `default:"false"`
		// 主页
		Index string `default:"index.html"`
		// 错误页面
		Error string `default:"error.html"`
	}
}
