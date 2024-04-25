package step

import (
	"context"

	"github.com/goexl/args"
	"github.com/pangum/drone/internal/internal/command"
	"github.com/pangum/drone/internal/internal/config"
)

type Test struct {
	golang  *command.Golang
	test    *config.Test
	project *config.Project

	defaultFlags []string
}

func NewTest(golang *command.Golang, test *config.Test, project *config.Project) *Test {
	return &Test{
		golang:  golang,
		test:    test,
		project: project,

		defaultFlags: []string{
			// 缩短长时间运行的测试的测试时间
			"-short",
			// 随机
			"-shuffle=on",
		},
	}
}

func (t *Test) Runnable() bool {
	return nil != t.test.Enabled && *t.test.Enabled
}

func (t *Test) Run(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("test")
	// 加入默认测试参数
	arguments.Add(t.flags()...)
	// 加入测试文件
	arguments.Add(t.project.Source)
	// 执行测试命令
	err = t.golang.Exec(ctx, arguments.Build())

	return
}

func (t *Test) flags() (flags []any) {
	flags = make([]any, 0)
	if t.golang.Default() {
		for _, flag := range t.defaultFlags {
			flags = append(flags, flag)
		}
	}
	for _, flag := range t.test.Flags {
		flags = append(flags, flag)
	}

	return
}
