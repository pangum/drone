package config

import (
	"path/filepath"

	"github.com/pangum/drone/internal/internal/core"
)

type Output struct {
	// 文件名
	Name string `default:"${DRONE_STAGE_NAME}" json:"name,omitempty"`
	// 操作系统
	Os string `default:"linux" json:"os,omitempty"`
	// 架构
	Arch string `default:"amd64" json:"arch,omitempty"`
	// 版本
	Arm int `default:"7" json:"arm,omitempty"`
	// 是否开启
	Cgo *bool `default:"${CGO=false}" json:"cgo,omitempty"`
	// 编译模式
	Mode core.Mode `default:"release" json:"mode,omitempty" validate:"oneof=release debug"`
	// 环境变量
	Environments map[string]string `default:"${ENVIRONMENTS}" json:"environments,omitempty"`
}

func (o *Output) Filename(project *Project) string {
	return filepath.Join(project.Dir, o.Name)
}
