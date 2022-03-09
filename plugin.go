package main

import (
	`fmt`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 输入文件
	Input string `default:"${PLUGIN_INPUT=${INPUT=.}}"`
	// 输出文件
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=${DRONE_STAGE_NAME}}}"`
	// 编译模式
	Mode mode `default:"${PLUGIN_MODE=${MODE=release}}" validate:"oneof=release debug"`
	// 环境变量
	Envs []string `default:"${PLUGIN_ENVS=${ENVS}}"`

	// 应用名称
	Name string `default:"${PLUGIN_NAME=${NAME=${DRONE_STAGE_NAME}}}"`
	// 应用版本
	Version string `default:"${PLUGIN_VERSION=${VERSION=${DRONE_TAG=${DRONE_COMMIT_BRANCH}}}}"`
	// 编译版本
	Build string `default:"${PLUGIN_BUILD=${BUILD=${DRONE_BUILD_NUMBER}}}"`
	// 编译时间
	Timestamp string `default:"${PLUGIN_TIMESTAMP=${TIMESTAMP=${DRONE_BUILD_STARTED}}}"`
	// 分支版本
	Revision string `default:"${PLUGIN_REVISION=${REVISION=${DRONE_COMMIT_SHA}}}"`
	// 分支
	Branch string `default:"${PLUGIN_BRANCH=${BRANCH=${DRONE_COMMIT_BRANCH}}}"`

	// 代码检查
	Lint lint `default:"${PLUGIN_LINT=${LINT}}"`
	// 测试
	Test test `default:"${PLUGIN_TEST=${TEST}}"`
	// 压缩
	Compress compress `default:"${PLUGIN_COMPRESS=${COMPRESS}}"`

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

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.tidy, drone.Name(`清理`)),
		drone.NewStep(p.lint, drone.Name(`检查`)),
		drone.NewStep(p.test, drone.Name(`测试`)),
		drone.NewStep(p.build, drone.Name(`编译`)),
		drone.NewStep(p.compress, drone.Name(`压缩`)),
	}
}

func (p *plugin) Setup() (unset bool, err error) {
	p.defaultEnvs = []string{
		`CGO_ENABLED=0`,
		`GOOS=linux`,
	}
	p.defaultLinters = []string{
		`goerr113`,
		`nlreturn`,
		`bodyclose`,
		`rowserrcheck`,
		`gosec`,
		`unconvert`,
		`misspell`,
		`lll`,
	}
	p.defaultFlags = []string{
		// 删除掉符号表
		`-s`,
		// 去掉调试信息，无法使用GDB调试程序
		`-w`,
	}
	p.defaultTestFlags = []string{
		// 缩短长时间运行的测试的测试时间
		`-short`,
		// 随机
		`-shuffle=on`,
	}

	return
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`input`, p.Input),
		field.String(`output`, p.Output),
		field.Any(`lint`, p.Lint),

		field.String(`name`, p.Name),
		field.String(`version`, p.Version),
		field.String(`build`, p.Build),
		field.String(`timestamp`, p.Timestamp),
		field.String(`revision`, p.Revision),
		field.String(`branch`, p.Branch),
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

func (p *plugin) testFlags() (flags []interface{}) {
	flags = make([]interface{}, 0)
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

func (p *plugin) flags() (flags []string) {
	flags = make([]string, 0)
	if p.Defaults && modeRelease == p.Mode {
		flags = append(flags, p.defaultFlags...)
	}
	if `` != p.Name {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Name=%s'`, p.Name))
	}
	if `` != p.Version {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Version=%s'`, p.Version))
	}
	if `` != p.Build {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Build=%s'`, p.Build))
	}
	if `` != p.Timestamp {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Timestamp=%s'`, p.Timestamp))
	}
	if `` != p.Revision {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Revision=%s'`, p.Revision))
	}
	if `` != p.Branch {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Branch=%s'`, p.Branch))
	}

	return
}
