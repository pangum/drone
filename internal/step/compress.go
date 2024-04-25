package step

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/guc"
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
	outputs []*config.Output, project *config.Project,
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
	wg := new(guc.WaitGroup)
	wg.Add(len(c.outputs))
	for _, out := range c.outputs {
		go c.run(ctx, out, wg, &err)
	}
	wg.Wait()

	return
}

func (c *Compress) run(ctx *context.Context, output *config.Output, wg *guc.WaitGroup, err *error) {
	defer wg.Done()

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
	command := c.base.Command(c.binary.Upx).Context(*ctx).Args(arguments.Build()).Dir(c.project.Source)
	if _, ce := command.Build().Exec(); nil != ce {
		*err = ce
	}
}
