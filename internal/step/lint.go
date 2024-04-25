package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/pangum/drone/internal/internal/config"

	"github.com/goexl/args"
)

type Lint struct {
	base    *drone.Base
	binary  *config.Binary
	lint    *config.Lint
	project *config.Project

	defaultLinters []string
}

func NewLint(base *drone.Base, binary *config.Binary, lint *config.Lint, project *config.Project) *Lint {
	return &Lint{
		base:    base,
		binary:  binary,
		lint:    lint,
		project: project,

		defaultLinters: []string{
			"goerr113",
			"nlreturn",
			"bodyclose",
			"rowserrcheck",
			"gosec",
			"unconvert",
			"misspell",
			"lll",
		},
	}
}

func (l *Lint) Runnable() bool {
	return nil != l.lint.Enabled && *l.lint.Enabled
}

func (l *Lint) Run(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("run")
	// 设置超时时间
	arguments.Argument("timeout", l.lint.Timeout)
	// 始终显示
	arguments.Argument("color", "always")
	// 显示详细信息
	if l.base.Verbose {
		arguments.Flag("verbose")
	}
	// 显示调试信息
	for _, linter := range l.linters() {
		arguments.Argument("enable", linter)
	}

	// 执行代码检查命令
	command := l.base.Command(l.binary.Lint).Context(*ctx).Args(arguments.Build()).Dir(l.project.Source)
	_, err = command.Build().Exec()

	return
}

func (l *Lint) linters() (linters []string) {
	linters = make([]string, 0)
	if l.base.Default() {
		linters = append(linters, l.defaultLinters...)
	}
	linters = append(linters, l.lint.Linters...)

	return
}
