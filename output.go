package main

import (
	"strings"
)

type output struct {
	// 文件名
	Name string `default:"${DRONE_STAGE_NAME}" json:"name"`
	// 操作系统
	Os string `default:"linux" json:"os"`
	// 架构
	Arch string `default:"amd64" json:"arch"`
	// 编译模式
	Mode mode `default:"release" json:"mode" validate:"oneof=release debug"`
	// 环境变量
	Envs []string `json:"envs"`
}

func (o *output) build(plugin *plugin) (err error) {
	args := []any{
		"build",
		"-o",
		o.Name,
	}
	if plugin.Verbose {
		args = append(args, "-x")
	}

	// 写入编译标签
	args = append(args, "-ldflags", strings.Join(plugin.flags(o.Mode), ` `))

	// 执行编译命令
	command := plugin.Command(goExe).Args(args...).Dir(plugin.Source)
	command.Env("GOOS", o.Os).Env("GOARCH", o.Arch)
	command.StringEnvs(plugin.envs()...)
	command.StringEnvs(o.Envs...)
	err = command.Exec()

	return
}
