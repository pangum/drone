package main

import (
	"fmt"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 源文件目录
	Source string `default:"${SOURCE=.}"`
	// 输出文件
	Output *output `default:"${OUTPUT}"`
	// 输出列表
	Outputs []*output `default:"${OUTPUTS}"`
	// 环境变量
	Envs []string `default:"${ENVS}"`

	// 应用名称
	Name string `default:"${NAME=${DRONE_STAGE_NAME}}"`
	// 应用版本
	Version string `default:"${VERSION=${DRONE_TAG=${DRONE_COMMIT_BRANCH}}}"`
	// 编译版本
	Build string `default:"${BUILD=${DRONE_BUILD_NUMBER}}"`
	// 编译时间
	Timestamp string `default:"${TIMESTAMP=${DRONE_BUILD_STARTED}}"`
	// 分支版本
	Revision string `default:"${REVISION=${DRONE_COMMIT_SHA}}"`
	// 分支
	Branch string `default:"${BRANCH=${DRONE_COMMIT_BRANCH}}"`

	// 代码检查
	Lint lint `default:"${LINT}"`
	// 测试
	Test test `default:"${TEST}"`
	// 压缩
	Compress compress `default:"${COMPRESS}"`

	defaultEnvs      []string
	defaultLinters   []string
	defaultFlags     []string
	defaultTestFlags []string
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.tidy, drone.Name("清理")),
		drone.NewStep(p.lint, drone.Name("检查")),
		drone.NewStep(p.test, drone.Name("测试")),
		drone.NewStep(p.build, drone.Name("编译")),
		drone.NewStep(p.compress, drone.Name("压缩")),
	}
}

func (p *plugin) Setup() (unset bool, err error) {
	if nil != p.Output {
		p.Outputs = append(p.Outputs, p.Output)
	}

	p.defaultEnvs = []string{
		"CGO_ENABLED=0",
	}
	p.defaultLinters = []string{
		"goerr113",
		"nlreturn",
		"bodyclose",
		"rowserrcheck",
		"gosec",
		"unconvert",
		"misspell",
		"lll",
	}
	p.defaultFlags = []string{
		// 删除掉符号表
		"-s",
		// 去掉调试信息，无法使用GDB调试程序
		"-w",
	}
	p.defaultTestFlags = []string{
		// 缩短长时间运行的测试的测试时间
		"-short",
		// 随机
		"-shuffle=on",
	}

	return
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("input", p.Source),
		field.New("output", p.Output),
		field.New("lint", p.Lint),

		field.New("name", p.Name),
		field.New("version", p.Version),
		field.New("build", p.Build),
		field.New("timestamp", p.Timestamp),
		field.New("revision", p.Revision),
		field.New("branch", p.Branch),
	}
}

func (p *plugin) linters() (linters []string) {
	linters = make([]string, 0)
	if p.Defaults {
		linters = append(linters, p.defaultLinters...)
	}
	linters = append(linters, p.Lint.Linters...)

	return
}

func (p *plugin) testFlags() (flags []any) {
	flags = make([]any, 0)
	if p.Defaults {
		for _, flag := range p.defaultTestFlags {
			flags = append(flags, flag)
		}
	}
	for _, flag := range p.Test.Flags {
		flags = append(flags, flag)
	}

	return
}

func (p *plugin) envs() (envs []string) {
	envs = make([]string, 0, len(p.Envs)+2)
	if p.Defaults {
		envs = append(envs, p.defaultEnvs...)
	}
	envs = append(envs, p.Envs...)

	return
}

func (p *plugin) flags(mode mode) (flags []string) {
	flags = make([]string, 0)
	if p.Defaults && modeRelease == mode {
		flags = append(flags, p.defaultFlags...)
	}
	if "" != p.Name {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Name=%s'", p.Name))
	}
	if "" != p.Version {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Version=%s'", p.Version))
	}
	if "" != p.Build {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Build=%s'", p.Build))
	}
	if "" != p.Timestamp {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Timestamp=%s'", p.Timestamp))
	}
	if "" != p.Revision {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Revision=%s'", p.Revision))
	}
	if "" != p.Branch {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu.Branch=%s'", p.Branch))
	}

	return
}
