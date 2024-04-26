package config

type Cos struct {
	// 存储桶地址
	Endpoint string `default:"${ENDPOINT}" validate:"required,url" json:"endpoint,omitempty"`

	// 分隔符
	Separator string `default:"${SEPARATOR=/}" json:"separator,omitempty"`
	// 是否清空存储桶
	Clear *bool `default:"${CLEAR=true}" json:"clear,omitempty"`
	// 路径前缀，所有文件上传都会在这上面加上前缀
	Prefix string `default:"${PREFIX}" json:"prefix,omitempty"`
	// 路径后缀，所有文件上传都会在这上面加上后缀
	Suffix string `default:"${SUFFIX}" json:"suffix,omitempty"`

	// 静态网站
	Website Website `default:"${WEBSITE}" json:"website,omitempty"`
}
