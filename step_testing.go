package main

import (
	"context"
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
	args := []any{
		"test",
	}
	// 加入默认测试参数
	args = append(args, t.testFlags()...)
	// 加入测试文件
	args = append(args, t.Source)
	// 执行测试命令
	err = t.Command(goExe).Args(args...).Dir(t.Source).StringEnvs(t.envs()...).Exec()

	return
}
