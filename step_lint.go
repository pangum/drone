package main

import (
	"context"
	"fmt"

	"github.com/goexl/gox/args"
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
	lintArgs := args.New().Build().Subcommand("run").Arg("timeout", l.Lint.Timeout).Arg("color", "always")
	// 显示详细信息
	if l.Verbose {
		lintArgs.Flag("verbose")
	}
	// 显示调试信息
	for _, linter := range l.linters() {
		lintArgs.Arg("enable", linter)
	}

	// 执行代码检查命令
	command := l.Command(l.Lint.Binary).Args(lintArgs.Build()).Dir(l.Source)
	environment := command.Environment()
	environment.String(l.envs()...)
	command = environment.Build()
	_, err = command.Build().Exec()
	fmt.Println("===================")
	fmt.Println(l.Source)

	return
}
