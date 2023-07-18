package internal

import (
	"context"

	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/plugin"
)

type Test struct {
	*plugin.Plugin
}

func NewTest(plugin *plugin.Plugin) *Test {
	return &Test{
		Plugin: plugin,
	}
}

func (t *Test) Runnable() bool {
	return t.TestEnabled()
}

func (t *Test) Run(_ context.Context) (err error) {
	testArgs := args.New().Build().Subcommand("test")
	// 加入默认测试参数
	testArgs.Add(t.TestFlags()...)
	// 加入测试文件
	testArgs.Add(t.Source)
	// 执行测试命令
	command := t.Command(t.Binary.Go).Args(testArgs.Build()).Dir(t.Source)
	environment := command.Environment()
	environment.String(t.Environments()...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
