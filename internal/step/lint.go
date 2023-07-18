package step

import (
	"context"

	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/plugin"
)

type Lint struct {
	*plugin.Plugin
}

func NewLint(plugin *plugin.Plugin) *Lint {
	return &Lint{
		Plugin: plugin,
	}
}

func (l *Lint) Runnable() bool {
	return l.LintEnabled()
}

func (l *Lint) Run(_ context.Context) (err error) {
	lintArgs := args.New().Build().Subcommand("run").Arg("timeout", l.Lint.Timeout).Arg("color", "always")
	// 显示详细信息
	if l.Verbose {
		lintArgs.Flag("verbose")
	}
	// 显示调试信息
	for _, linter := range l.Linters() {
		lintArgs.Arg("enable", linter)
	}

	// 执行代码检查命令
	command := l.Command(l.Lint.Binary).Args(lintArgs.Build()).Dir(l.Source)
	environment := command.Environment()
	environment.String(l.Environments()...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
