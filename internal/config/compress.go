package config

import (
	"fmt"
	"strconv"

	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
	"github.com/pangum/drone/internal/plugin"
)

type Compress struct {
	// 启用压缩
	Enabled *bool `default:"true" json:"enabled"`
	// 类型
	Type core.CompressType `default:"upx" json:"type" validate:"oneof=upx"`
	// 压缩等级
	Level string `default:"lzma" json:"level" validate:"oneof=1 2 3 4 5 6 7 8 9 best lzma brute ultra-brute"`
}

func (c *Compress) Do(plugin *plugin.Plugin, output *Output) (err error) {
	switch c.Type {
	case core.CompressTypeUpx:
		err = c.upx(plugin, output)
	}

	return
}

func (c *Compress) upx(plugin *plugin.Plugin, output *Output) (err error) {
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
	upxArgs.Add(output.name(plugin))

	// 执行清理依赖命令
	command := plugin.Command(plugin.Binary.Upx).Args(upxArgs.Build()).Dir(plugin.Source)
	command = command.Environment().String(plugin.Environments()...).Build()
	_, err = command.Build().Exec()

	return
}
