package step

import (
	"context"
	"github.com/dronestock/drone"
	"github.com/pangum/drone/internal/config"
	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gox/args"
)

type Test struct {
	drone.Base
	internal.Core

	config  *config.Test
	outputs []*config.Output
	flags   []any
	envs    []string
}

func NewTest(
	base drone.Base, core internal.Core,
	config *config.Test, outputs []*config.Output, flags []any, envs []string,
) *Test {
	return &Test{
		Base: base,
		Core: core,

		config:  config,
		outputs: outputs,
		flags:   flags,
		envs:    envs,
	}
}

func (t *Test) Runnable() bool {
	return nil != t.config.Enabled && *t.config.Enabled
}

func (t *Test) Run(_ context.Context) (err error) {
	testArgs := args.New().Build().Subcommand("test")
	// 加入默认测试参数
	testArgs.Add(t.flags...)
	// 加入测试文件
	testArgs.Add(t.Source)
	// 执行测试命令
	command := t.Command(t.Binary.Go).Args(testArgs.Build()).Dir(t.Source)
	environment := command.Environment()
	environment.String(t.envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
