package step

import (
	"context"
	"fmt"
	"strings"

	"github.com/goexl/args"
	"github.com/goexl/gox"
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

func NewBuild(
	golang *command.Golang,
	outputs gox.Slice[*config.Output], project *config.Project,
	debug *config.Debug,
) *Build {
	return &Build{
		golang: golang,

		outputs: outputs,
		project: project,
		debug:   debug,

		defaultFlags: []string{
			"-s", // 删除掉符号表
			"-w", // 去掉调试信息，无法使用GDB调试程序
		},
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx *context.Context) (err error) {
	waiter := guc.New().Wait().Group(len(b.outputs))
	for _, output := range b.outputs {
		cloned := output
		go b.run(ctx, cloned, waiter, &err)
	}
	waiter.Wait()

	return
}

func (b *Build) run(ctx *context.Context, output *config.Output, waiter guc.Waiter, err *error) {
	defer waiter.Done()

	arguments := args.New().Long(constant.Strike).Build().Subcommand("build")
	arguments = arguments.Flag("o").Add(output.Filename(b.project))
	// 写入编译标签
	arguments.Argument("ldflags", strings.Join(b.flags(output.Mode), constant.Space))

	// 准备编译环境变量
	environments := []*core.Environment{
		core.NewEnvironment(constant.GoOS, output.Os),
		core.NewEnvironment(constant.GoArch, output.Arch),
	}
	if 0 != output.Arm && strings.Contains(output.Arch, "arm") {
		environments = append(environments, core.NewEnvironment(constant.GoArm, output.Arm))
	}
	if *output.Cgo {
		b.cgo(output, &environments)
	}
	for key, value := range output.Environments {
		environments = append(environments, core.NewEnvironment(key, value))
	}

	// 执行编译命令
	if be := b.golang.Exec(ctx, arguments.Build(), environments...); nil != be {
		*err = be
		b.golang.Warn("编译出错", field.New("output", output))
	}
}

func (b *Build) cgo(output *config.Output, environments *[]*core.Environment) {
	*environments = append(*environments, core.NewEnvironment(constant.Cgo, 1))
	if 7 == output.Arm {
		*environments = append(*environments, core.NewEnvironment(constant.CC, "arm-linux-gnueabihf-gcc"))
	}
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
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Build=%s'", b.debug.Build))
	}
	if "" != b.debug.Complied {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Compiled=%s'", b.debug.Complied))
	}
	if "" != b.debug.Revision {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Revision=%s'", b.debug.Revision))
	}
	if "" != b.debug.Branch {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Branch=%s'", b.debug.Branch))
	}

	return
}
