package config

import (
	"path/filepath"
	"strings"

	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
	"github.com/pangum/drone/internal/plugin"
)

type Output struct {
	// 文件名
	Name string `default:"${OUTPUT_NAME=${DRONE_STAGE_NAME}}" json:"name"`
	// 操作系统
	Os string `default:"${OUTPUT_OS=linux}" json:"os"`
	// 架构
	Arch string `default:"${OUTPUT_ARCH=amd64}" json:"arch"`
	// 编译模式
	Mode core.Mode `default:"${OUTPUT_MODE=release}" json:"mode" validate:"oneof=release debug"`
	// 环境变量
	Envs []string `default:"${OUTPUT_ENVS}" json:"envs"`
}

func (o *Output) Build(plugin *plugin.Plugin) (err error) {
	buildArgs := args.New().Long(core.Strike).Build().Subcommand("Build").Flag("o").Add(o.name(plugin))
	if plugin.Verbose {
		buildArgs.Flag("x")
	}

	// 写入编译标签
	buildArgs.Arg("ldflags", strings.Join(plugin.Flags(o.Mode), core.Space))

	// 执行编译命令
	command := plugin.Command(plugin.Binary.Go).Args(buildArgs.Build()).Dir(plugin.Source)
	environment := command.Environment()
	environment.Kv(core.Goos, o.Os)
	environment.Kv(core.Goarch, o.Arch)
	environment.String(plugin.Environments()...)
	environment.String(o.Envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}

func (o *Output) name(plugin *plugin.Plugin) string {
	return filepath.Join(plugin.Dir, o.Name)
}
