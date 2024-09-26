package step

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/drone/internal/internal/config"
)

type Compress struct {
	base     *drone.Base
	binary   *config.Binary
	compress *config.Compress
	project  *config.Project
	outputs  []*config.Output
}

func NewCompress(
	base *drone.Base, binary *config.Binary,
	compress *config.Compress,
	outputs gox.Slice[*config.Output], project *config.Project,
) *Compress {
	return &Compress{
		base:   base,
		binary: binary,

		compress: compress,
		outputs:  outputs,
		project:  project,
	}
}

func (c *Compress) Runnable() bool {
	return nil != c.compress.Enabled && *c.compress.Enabled
}

func (c *Compress) Run(ctx *context.Context) (err error) {
	for _, output := range c.outputs {
		if re := c.run(ctx, output); nil != re {
			err = re
			c.base.Warn("压缩程序出错", field.New("output", output), field.Error(err))
		} else {
			c.base.Info("压缩程序成功", field.New("output", output))
		}
	}

	return
}

func (c *Compress) run(ctx *context.Context, output *config.Output) (err error) {
	arguments := args.New().Build().Flag("mono").Flag("color").Flag("f").Flag("force-macos")
	if c.base.Verbose {
		arguments.Flag("v")
	}

	// 压缩等级
	if _, ce := strconv.Atoi(c.compress.Level); nil != ce {
		arguments.Add(fmt.Sprintf("--%s", c.compress.Level))
	} else {
		arguments.Add(fmt.Sprintf("-%s", c.compress.Level))
	}
	// 添加输出文件
	arguments.Add(output.Filename(c.project))

	// 执行清理依赖命令
	c.base.Info("压缩程序开始", field.New("output", output))
	command := c.base.Command(c.binary.Upx).Context(*ctx).Arguments(arguments.Build()).Dir(c.project.Source)
	_, err = command.Build().Exec()

	return
}
