package plugin

import (
	"fmt"

	"github.com/pangum/drone/internal/plugin/internal"
	"github.com/pangum/drone/internal/plugin/internal/step"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/drone/internal/config"
	"github.com/pangum/drone/internal/core"
)

type Plugin struct {
	internal.Core

	// 输出文件
	Output *config.Output `default:"${OUTPUT}" json:"output"`
	// 输出列表
	Outputs []*config.Output `default:"${OUTPUTS}" json:"outputs"`
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
	Complied string `default:"${TIMESTAMP=${DRONE_BUILD_STARTED}}" json:"timestamp"`
	// 分支版本
	Revision string `default:"${REVISION=${DRONE_COMMIT_SHA}}" json:"revision"`
	// 分支
	Branch string `default:"${BRANCH=${DRONE_COMMIT_BRANCH}}" json:"branch"`

	// 内存对齐
	Alignment config.Alignment `default:"${ALIGNMENT}" json:"alignment,omitempty"`
	// 代码检查
	Lint config.Lint `default:"${LINT}" json:"lint"`
	// 测试
	Test config.Test `default:"${TEST}" json:"test"`
	// 压缩
	Compress config.Compress `default:"${COMPRESS}" json:"compress"`

	defaultEnvs      []string
	defaultLinters   []string
	defaultFlags     []string
	defaultTestFlags []string
}

func New() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewTidy(&p.Core, p.Environments())).Name("依赖清理").Build(),
		drone.NewStep(step.NewAlignment(&p.Core, &p.Alignment, p.Environments())).Name("内存对齐").Build(),
		drone.NewStep(step.NewLint(&p.Core, &p.Lint, p.Linters(), p.Environments())).Name("静态检查").Build(),
		drone.NewStep(step.NewTest(&p.Core, &p.Test, p.Outputs, p.TestFlags(), p.Environments())).Name("单元测试").Build(),
		drone.NewStep(step.NewBuild(&p.Core, p.Outputs, p.Flags, p.Environments())).Name("编译打包").Break().Build(),
		drone.NewStep(step.NewCompress(&p.Core, &p.Compress, p.Outputs, p.Environments())).Name("程序压缩").Build(),
	}
}

func (p *Plugin) Setup() (err error) {
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

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("output", p.Output),
		field.New("lint", p.Lint),

		field.New("name", p.Name),
		field.New("version", p.Version),
		field.New("build", p.Build),
		field.New("complied", p.Complied),
		field.New("revision", p.Revision),
		field.New("branch", p.Branch),
	}
}

func (p *Plugin) Linters() (linters []string) {
	linters = make([]string, 0)
	if p.Default() {
		linters = append(linters, p.defaultLinters...)
	}
	linters = append(linters, p.Lint.Linters...)

	return
}

func (p *Plugin) TestFlags() (flags []any) {
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

func (p *Plugin) Environments() (envs []string) {
	envs = make([]string, 0, len(p.Envs)+2)
	if p.Default() {
		envs = append(envs, p.defaultEnvs...)
	}
	for _, private := range p.Privates {
		_goPrivate := gox.StringBuilder(core.GoPrivate, core.Equal, private).String()
		_goNoProxy := gox.StringBuilder(core.GoNoProxy, core.Equal, private).String()
		envs = append(envs, _goPrivate, _goNoProxy)
	}
	envs = append(envs, p.Envs...)

	return
}

func (p *Plugin) Flags(mode core.Mode) (flags []string) {
	flags = make([]string, 0)
	if p.Default() && core.ModeRelease == mode {
		flags = append(flags, p.defaultFlags...)
	}
	if "" != p.Name {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Name=%s'", p.Name))
	}
	if "" != p.Version {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Version=%s'", p.Version))
	}
	if "" != p.Build {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Build=%s'", p.Build))
	}
	if "" != p.Complied {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Complied=%s'", p.Complied))
	}
	if "" != p.Revision {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Revision=%s'", p.Revision))
	}
	if "" != p.Branch {
		flags = append(flags, fmt.Sprintf("-X 'github.com/pangum/pangu/internal.Branch=%s'", p.Branch))
	}

	return
}
