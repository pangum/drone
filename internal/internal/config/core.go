package config

import (
	"github.com/dronestock/drone"
)

type Core struct {
	drone.Base

	// 控制程序
	Binary Binary `default:"${BINARY}" json:"binary,omitempty"`
	// 源文件目录
	Source string `default:"${SOURCE=.}" json:"source,omitempty"`
	// 输出目录
	Dir string `default:"${DIR=.}" json:"dir,omitempty"`
}
