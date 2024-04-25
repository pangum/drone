package internal

import (
	"github.com/pangum/drone/internal/internal/command"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/step"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 控制程序
	Binary config.Binary `default:"${BINARY}" json:"binary,omitempty"`
	// 项目信息
	Project config.Project `default:"${PROJECT}" json:"project,omitempty"`
	// 输出文件
	Output *config.Output `default:"${OUTPUT}" json:"output,omitempty"`
	// 输出列表
	Outputs []*config.Output `default:"${OUTPUTS}" json:"outputs,omitempty"`
	// 调试信息
	Debug config.Debug `default:"${DEBUG}" json:"debug,omitempty"`
	// 内存对齐
	Alignment config.Alignment `default:"${ALIGNMENT}" json:"alignment,omitempty"`
	// 代码检查
	Lint config.Lint `default:"${LINT}" json:"lint,omitempty"`
	// 测试
	Test config.Test `default:"${TEST}" json:"test,omitempty"`
	// 压缩
	Compress config.Compress `default:"${COMPRESS}" json:"compress,omitempty"`

	golang *command.Golang
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewTidy(p.golang, &p.Project)).Name("依赖清理").Build(),
		// nolint:lll
		drone.NewStep(step.NewAlignment(&p.Base, &p.Binary, &p.Alignment, &p.Project)).Name("内存对齐").Interrupt().Continue().Build(),
		drone.NewStep(step.NewLint(&p.Base, &p.Binary, &p.Lint, &p.Project)).Name("静态检查").Build(),
		drone.NewStep(step.NewTest(p.golang, &p.Test, &p.Project)).Name("单元测试").Build(),
		drone.NewStep(step.NewBuild(p.golang, p.Outputs, &p.Project, &p.Debug)).Name("编译打包").Break().Build(),
		drone.NewStep(step.NewCompress(&p.Base, &p.Binary, &p.Compress, p.Outputs, &p.Project)).Name("程序压缩").Build(),
	}
}

func (p *plugin) Setup() (err error) {
	if nil != p.Output {
		p.Outputs = append(p.Outputs, p.Output)
	}
	p.golang = command.NewGolang(&p.Base, &p.Binary, &p.Project)

	return
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("binary", p.Binary),
		field.New("project", p.Project),
		field.New("debug", p.Debug),
		field.New("alignment", p.Alignment),
		field.New("lint", p.Lint),
		field.New("test", p.Test),
		field.New("compress", p.Compress),
	}
}
