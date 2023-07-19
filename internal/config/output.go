package config

import (
	"path/filepath"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
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

func (o *Output) Build(
	plugin *drone.Base,
	binary *Binary,
	source string, dir string,
	flags []string, envs []string,
) (err error) {
	buildArgs := args.New().Long(core.Strike).Build().Subcommand("build").Flag("o").Add(o.name(dir))
	if plugin.Verbose {
		buildArgs.Flag("x")
	}

	// 写入编译标签
	buildArgs.Arg("ldflags", strings.Join(flags, core.Space))

	// 执行编译命令
	command := plugin.Command(binary.Go).Args(buildArgs.Build()).Dir(source)
	environment := command.Environment()
	environment.Kv(core.Goos, o.Os)
	environment.Kv(core.Goarch, o.Arch)
	environment.String(envs...)
	environment.String(o.Envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}

func (o *Output) name(dir string) string {
	return filepath.Join(dir, o.Name)
}
