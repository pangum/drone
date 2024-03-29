package config

import (
	"fmt"
	"strconv"

	"github.com/dronestock/drone"
	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
)

type Compress struct {
	// 启用压缩
	Enabled *bool `default:"true" json:"enabled"`
	// 类型
	Type core.CompressType `default:"upx" json:"type" validate:"oneof=upx"`
	// 压缩等级
	Level string `default:"lzma" json:"level" validate:"oneof=1 2 3 4 5 6 7 8 9 best lzma brute ultra-brute"`
}

func (c *Compress) Do(
	base *drone.Base,
	binary *Binary,
	source string, dir string, verbose bool,
	output *Output, envs []string,
) (err error) {
	switch c.Type {
	case core.CompressTypeUpx:
		err = c.upx(base, binary, source, dir, verbose, output, envs)
	}

	return
}

func (c *Compress) upx(
	base *drone.Base,
	binary *Binary,
	source string, dir string, verbose bool,
	output *Output, envs []string,
) (err error) {
	upxArgs := args.New().Build().Flag("mono").Flag("color").Flag("f").Flag("force-macos")
	if verbose {
		upxArgs.Flag("v")
	}

	// 压缩等级
	if _, ce := strconv.Atoi(c.Level); nil != ce {
		upxArgs.Add(fmt.Sprintf("--%s", c.Level))
	} else {
		upxArgs.Add(fmt.Sprintf("-%s", c.Level))
	}
	// 添加输出文件
	upxArgs.Add(output.name(dir))

	// 执行清理依赖命令
	command := base.Command(binary.Upx).Args(upxArgs.Build()).Dir(source)
	command = command.Environment().String(envs...).Build()
	_, err = command.Build().Exec()

	return
}
