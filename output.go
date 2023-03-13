package main

import (
	"path/filepath"
	"strings"

	"github.com/goexl/gox/args"
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
	buildArgs := args.New().Long(strike).Build().Subcommand("build").Flag("o").Add(o.name(plugin))
	if plugin.Verbose {
		buildArgs.Flag("x")
	}

	// 写入编译标签
	buildArgs.Arg("ldflags", strings.Join(plugin.flags(o.Mode), space))

	// 执行编译命令
	command := plugin.Command(goExe).Args(buildArgs.Build()).Dir(plugin.Source)
	environment := command.Environment()
	environment.Kv(goos, o.Os)
	environment.Kv(goarch, o.Arch)
	environment.String(plugin.envs()...)
	environment.String(o.Envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}

func (o *output) name(plugin *plugin) string {
	return filepath.Join(plugin.Dir, o.Name)
}
