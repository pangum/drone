package main

import (
	"strings"

	"github.com/dronestock/drone"
)

type output struct {
	// 文件名
	Name string `default:"${DRONE_STAGE_NAME}" json:"name"`
	// 操作系统
	Os string `default:"linux" json:"os"`
	// 架构
	Arch string `default:"amd64" json:"arch"`
	// 环境变量
	Envs []string `json:"envs"`
}

func (o *output) build(plugin *plugin) (err error) {
	args := []interface{}{
		`build`,
		`-o`,
		o.Name,
	}
	if plugin.Verbose {
		args = append(args, `-x`)
	}

	// 写入编译标签
	args = append(args, `-ldflags`, strings.Join(plugin.flags(), ` `))

	// 执行编译命令
	options := drone.NewExecOptions(
		drone.Args(args...),
		drone.Dir(plugin.Source),
		drone.Env("GOOS", o.Os), drone.Env("GOARCH", o.Arch),
	)
	options = append(options, drone.StringEnvs(plugin.Envs...))
	options = append(options, drone.StringEnvs(o.Envs...))
	err = plugin.Exec(goExe, options...)

	return
}
