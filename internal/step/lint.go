package step

import (
	"context"

	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gox/args"
)

type Lint struct {
	*internal.Core

	config  *config.Lint
	envs    []string
	linters []string
}

func NewLint(core *internal.Core, config *config.Lint, linters []string, envs []string) *Lint {
	return &Lint{
		Core: core,

		config:  config,
		linters: linters,
		envs:    envs,
	}
}

func (l *Lint) Runnable() bool {
	return nil != l.config.Enabled && *l.config.Enabled
}

func (l *Lint) Run(_ context.Context) (err error) {
	lintArgs := args.New().Build().Subcommand("run").Arg("timeout", l.config.Timeout).Arg("color", "always")
	// 显示详细信息
	if l.Verbose {
		lintArgs.Flag("verbose")
	}
	// 显示调试信息
	for _, linter := range l.linters {
		lintArgs.Arg("enable", linter)
	}

	// 执行代码检查命令
	command := l.Command(l.Binary.Lint).Args(lintArgs.Build()).Dir(l.Source)
	environment := command.Environment()
	environment.String(l.envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
