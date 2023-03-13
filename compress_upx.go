package main

import (
	"fmt"
	"strconv"

	"github.com/goexl/gox/args"
)

func (c *compress) upx(plugin *plugin, output *output) (err error) {
	upxArgs := args.New().Build().Flag("mono").Flag("color").Flag("f")
	if plugin.Verbose {
		upxArgs.Flag("v")
	}

	// 压缩等级
	if _, ce := strconv.Atoi(c.Level); nil != ce {
		upxArgs.Add(fmt.Sprintf("--%s", c.Level))
	} else {
		upxArgs.Add(fmt.Sprintf("-%s", c.Level))
	}
	// 添加输出文件
	upxArgs.Add(output.Name)

	// 执行清理依赖命令
	command := plugin.Command(upxExe).
		Args(upxArgs.Build()).
		Dir(plugin.Source).
		StringEnvironment(plugin.envs()...)
	_, err = command.Build().Exec()

	return
}
