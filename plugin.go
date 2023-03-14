package main

import (
	"fmt"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 控制程序
	Binary string `default:"${LINT_BINARY=go}" json:"binary"`
	// 源文件目录
	Source string `default:"${SOURCE=.}" json:"source"`
	// 输出目录
	Dir string `default:"${DIR=.}" json:"dir"`
	// 输出文件
	Output *output `default:"${OUTPUT}" json:"output"`
	// 输出列表
	Outputs []*output `default:"${OUTPUTS}" json:"outputs"`
	// 私有库
	Privates []string `default:"${PRIVATES}" json:"privates"`
	// 环境变量
	Envs []string `default:"${ENVS}" json:"envs"`

	// 应用名称
	Name string `default:"${NAME=${DRONE_STAGE_NAME}}" json:"name"`
	// 应用版本
	Version string `default:"${VERSION=${DRONE_TAG=${DRONE_COMMIT_BRANCH}}}" json:"version"`
	// 编译版本
	Build string `default:"${BUILD=${DRONE_BUILD_NUMBER}}" json:"build"`
	// 编译时间
	Timestamp string `default:"${TIMESTAMP=${DRONE_BUILD_STARTED}}" json:"timestamp"`
	// 分支版本
	Revision string `default:"${REVISION=${DRONE_COMMIT_SHA}}" json:"revision"`
	// 分支
	Branch string `default:"${BRANCH=${DRONE_COMMIT_BRANCH}}" json:"branch"`

	// 代码检查
	Lint lint `default:"${LINT}" json:"lint"`
	// 测试
	Test test `default:"${TEST}" json:"test"`
	// 压缩
	Compress compress `default:"${COMPRESS}" json:"compress"`

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
		drone.NewStep(newTidyStep(p)).Name("清理").Build(),
		drone.NewStep(newLintStep(p)).Name("检查").Build(),
		drone.NewStep(newTestStep(p)).Name("测试").Build(),
		drone.NewStep(newBuildStep(p)).Name("编译").Build(),
		drone.NewStep(newCompressStep(p)).Name("压缩").Build(),
	}
}

func (p *plugin) Setup() (err error) {
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
		field.New("source", p.Source),
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
	if p.Default() {
		linters = append(linters, p.defaultLinters...)
	}
	linters = append(linters, p.Lint.Linters...)

	return
}

func (p *plugin) testFlags() (flags []any) {
	flags = make([]any, 0)
	if p.Default() {
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
	if p.Default() {
		envs = append(envs, p.defaultEnvs...)
	}
	for _, private := range p.Privates {
		_goPrivate := gox.StringBuilder(goPrivate, equal, private).String()
		_goNoProxy := gox.StringBuilder(goNoProxy, equal, private).String()
		envs = append(envs, _goPrivate, _goNoProxy)
	}
	envs = append(envs, p.Envs...)

	return
}

func (p *plugin) flags(mode mode) (flags []string) {
	flags = make([]string, 0)
	if p.Default() && modeRelease == mode {
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

func (p *plugin) testEnabled() bool {
	return nil != p.Test.Enabled && *p.Test.Enabled
}

func (p *plugin) lintEnabled() bool {
	return nil != p.Lint.Enabled && *p.Lint.Enabled
}

func (p *plugin) compressEnabled() bool {
	return nil != p.Compress.Enabled && *p.Compress.Enabled
}
