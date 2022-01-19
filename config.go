package main

import (
	`fmt`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 输入文件
	Input string `default:"${PLUGIN_INPUT=${INPUT=.}}"`
	// 输出文件
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=${DRONE_STAGE_NAME}}}"`
	// 环境变量
	Envs []string `default:"${PLUGIN_ENVS=${ENVS}}"`

	// 是否启用Lint插件
	Lint bool `default:"${PLUGIN_LINT=${LINT=true}}"`
	// 启用的Linter
	Linters []string `default:"${PLUGIN_LINTERS=${LINTERS}}"`

	// 是否启用测试
	Test bool `default:"${PLUGIN_TEST=${TEST=true}}"`

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

	defaultEnvs    []string
	defaultLinters []string
	defaultFlags   []string
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`input`, c.Input),
		field.String(`output`, c.Output),
		field.Bool(`lint`, c.Lint),

		field.String(`name`, c.Name),
		field.String(`version`, c.Version),
		field.String(`build`, c.Build),
		field.String(`timestamp`, c.Timestamp),
		field.String(`revision`, c.Revision),
		field.String(`branch`, c.Branch),
	}
}

func (c *config) linters() (linters []string) {
	linters = make([]string, 0)
	if c.Defaults {
		linters = append(linters, c.defaultLinters...)
	}
	linters = append(linters, c.Linters...)

	return
}

func (c *config) flags() (flags []string) {
	flags = make([]string, 0)
	if c.Defaults {
		flags = append(flags, c.defaultFlags...)
	}
	if `` != c.Name {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Name=%s'`, c.Name))
	}
	if `` != c.Version {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Version=%s'`, c.Version))
	}
	if `` != c.Build {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Build=%s'`, c.Build))
	}
	if `` != c.Timestamp {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Timestamp=%s'`, c.Timestamp))
	}
	if `` != c.Revision {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Revision=%s'`, c.Revision))
	}
	if `` != c.Branch {
		flags = append(flags, fmt.Sprintf(`-X 'github.com/pangum/pangu.Branch=%s'`, c.Branch))
	}

	return
}
