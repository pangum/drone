package internal

import "github.com/pangum/drone/internal/config"

type Core struct {
	// 控制程序
	Binary config.Binary `default:"${BINARY}" json:"binary,omitempty"`
	// 源文件目录
	Source string `default:"${SOURCE=.}" json:"source"`
	// 输出目录
	Dir string `default:"${DIR=.}" json:"dir"`
}
