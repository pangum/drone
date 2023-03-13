package main

import (
	"context"

	"github.com/goexl/gox/args"
)

type stepTest struct {
	*plugin
}

func newTestStep(plugin *plugin) *stepTest {
	return &stepTest{
		plugin: plugin,
	}
}

func (t *stepTest) Runnable() bool {
	return t.testEnabled()
}

func (t *stepTest) Run(_ context.Context) (err error) {
	testArgs := args.New().Build().Subcommand("test")
	// 加入默认测试参数
	testArgs.Add(t.testFlags()...)
	// 加入测试文件
	testArgs.Add(t.Source)
	// 执行测试命令
	_, err = t.Command(goExe).Args(testArgs.Build()).Dir(t.Source).StringEnvironment(t.envs()...).Build().Exec()

	return
}
