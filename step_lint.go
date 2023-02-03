package main

import (
	"context"
)

type stepLint struct {
	*plugin
}

func newLintStep(plugin *plugin) *stepLint {
	return &stepLint{
		plugin: plugin,
	}
}

func (l *stepLint) Runnable() bool {
	return l.lintEnabled()
}

func (l *stepLint) Run(_ context.Context) (err error) {
	args := []any{
		"run",
		"--timeout",
		l.Lint.Timeout,
		"--color",
		"always",
	}
	// 显示详细信息
	if l.Verbose {
		args = append(args, "--verbose")
	}
	// 显示调试信息
	for _, linter := range l.linters() {
		args = append(args, "--enable", linter)
	}

	// 执行代码检查命令
	err = l.Command(lintExe).Args(args...).Dir(l.Source).StringEnvs(l.envs()...).Exec()

	return
}
