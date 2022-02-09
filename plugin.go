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

	// 是否启用Lint插件
	Lint bool `default:"${PLUGIN_LINT=${LINT=true}}"`
	// 启用的Linter
	Linters []string `default:"${PLUGIN_LINTERS=${LINTERS}}"`

	// 是否启用测试
	Test bool `default:"${PLUGIN_TEST=${TEST=true}}"`
	// 测试参数
	TestArgs []string `default:"${PLUGIN_TEST_ARGS=${TEST_ARGS}}"`
	// 测试标志
	TestFlags []string `default:"${PLUGIN_TEST_FLAGS=${TEST_FLAGS}}"`

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

	// 启用压缩
	Upx bool `default:"${PLUGIN_UPX=${UPX=true}}"`
	// 压缩等级
	// nolint:lll
	UpxLevel string `default:"${PLUGIN_UPX_LEVEL=${UPX_LEVEL=ultra-brute}}" validate:"oneof=1 2 3 4 5 6 7 8 9 ultra-brute brute"`

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
		drone.NewStep(p.tidy, drone.Name(`清理依赖`)),
		drone.NewStep(p.lint, drone.Name(`代码静态检查`)),
		drone.NewStep(p.test, drone.Name(`测试`)),
		drone.NewStep(p.build, drone.Name(`编译`)),
		drone.NewStep(p.upx, drone.Name(`压缩`)),
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
		field.Bool(`lint`, p.Lint),

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
	linters = append(linters, p.Linters...)

	return
}

func (p *plugin) testFlags() (flags []string) {
	flags = make([]string, 0)
	if p.Defaults {
		flags = append(flags, p.defaultTestFlags...)
	}
	flags = append(flags, p.TestFlags...)

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
