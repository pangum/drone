package main

import (
	"fmt"
	"strconv"

	"github.com/dronestock/drone"
)

func (c *compress) upx(plugin *plugin, output *output) (err error) {
	args := []interface{}{
		`--mono`,
		`--color`,
		`-f`,
	}
	if plugin.Verbose {
		args = append(args, `-v`)
	}

	// 压缩等级
	if _, convErr := strconv.Atoi(c.Level); nil != convErr {
		args = append(args, fmt.Sprintf(`--%s`, c.Level))
	} else {
		args = append(args, fmt.Sprintf(`-%s`, c.Level))
	}
	// 添加输出文件
	args = append(args, output.Name)

	// 执行清理依赖命令
	err = plugin.Exec(upxExe, drone.Args(args...), drone.Dir(plugin.Source), drone.StringEnvs(plugin.Envs...))

	return
}