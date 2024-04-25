package step

import (
	"context"
	"fmt"
	"strings"

	"github.com/goexl/args"
	"github.com/goexl/guc"
	"github.com/pangum/drone/internal/internal/command"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/constant"
	"github.com/pangum/drone/internal/internal/core"

	"github.com/goexl/gox/field"
)

type Build struct {
	golang *command.Golang

	outputs []*config.Output
	project *config.Project
	debug   *config.Debug

	defaultFlags []string
}

func NewBuild(golang *command.Golang, outputs []*config.Output, project *config.Project, debug *config.Debug) *Build {
	return &Build{
		golang: golang,

		outputs: outputs,
		project: project,
		debug:   debug,

		defaultFlags: []string{
			// 删除掉符号表
			"-s",
			// 去掉调试信息，无法使用GDB调试程序
			"-w",
		},
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx *context.Context) (err error) {
	wg := new(guc.WaitGroup)
	wg.Add(len(b.outputs))
	for _, out := range b.outputs {
		go b.run(ctx, out, wg, &err)
	}
	wg.Wait()

	return
}

func (b *Build) run(ctx *context.Context, output *config.Output, wg *guc.WaitGroup, err *error) {
	defer wg.Done()

	arguments := args.New().Long(constant.Strike).Build().Subcommand("build").Flag("o").Add(output.Filename(b.project))
	// 写入编译标签
	arguments.Argument("ldflags", strings.Join(b.flags(output.Mode), constant.Space))
	// 执行编译命令
	if be := b.golang.Exec(
		ctx, arguments.Build(),
		core.NewEnvironment(constant.Goos, output.Os), core.NewEnvironment(constant.Goarch, output.Arch),
	); nil != be {
		*err = be
		b.golang.Warn("编译出错", field.New("output", output))
	}

	return
}

func (b *Build) flags(mode core.Mode) (flags []string) {
	flags = make([]string, 0)
	if b.golang.Default() && core.ModeRelease == mode {
		flags = append(flags, b.defaultFlags...)
	}
	if "" != b.debug.Name {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Name=%s'", b.debug.Name))
	}
	if "" != b.debug.Version {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Version=%s'", b.debug.Version))
	}
	if "" != b.debug.Build {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.build=%s'", b.debug.Build))
	}
	if "" != b.debug.Complied {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Complied=%s'", b.debug.Complied))
	}
	if "" != b.debug.Revision {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Revision=%s'", b.debug.Revision))
	}
	if "" != b.debug.Branch {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Branch=%s'", b.debug.Branch))
	}

	return
}
