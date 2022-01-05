package main

import (
	"strconv"
	"time"

	"github.com/storezhang/gox"
	"github.com/storezhang/gox/field"
	"github.com/storezhang/mengpo"
)

type config struct {
	// 输入文件
	Input string `default:"${PLUGIN_INPUT=${INPUT=.}}"`
	// 输出文件
	Output string `default:"${PLUGIN_OUTPUT=${OUTPUT=${DRONE_STAGE_NAME}}}"`
	// 环境变量
	Envs []string `default:"${PLUGIN_ENVS=${ENVS}}"`
	// 是否启用默认配置
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示调试信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`

	// 是否启用Lint插件
	Lint bool `default:"${PLUGIN_LINT=${LINT=true}}"`
	// 启用的Linter
	Linters []string `default:"${PLUGIN_LINTERS=${LINTERS}}"`

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

	// 默认环境变量
	DefaultEnvs []string `default:"['CGO_ENABLED=0','GOOS=linux']"`
	// 默认启用的Linter列表
	// nolint:lll
	DefaultLinters []string `default:"['goerr113','nlreturn','bodyclose','rowserrcheck','gosec','unconvert','misspell','lll']"`
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

func (c *config) load() (err error) {
	// 处理环境变量为字符串的时候和默认值格式不兼容
	if err = parseEnvs(`ENVS`, `LINTERS`); nil != err {
		return
	}
	if err = mengpo.Set(c); nil != err {
		return
	}

	// 启用默认值
	if c.Defaults {
		c.Envs = append(c.Envs, c.DefaultEnvs...)
		c.Linters = append(c.Linters, c.DefaultLinters...)
	}

	// 将时间变换成易读形式
	if timestamp, parseErr := strconv.ParseInt(c.Timestamp, 10, 64); nil == parseErr {
		c.Timestamp = time.Unix(timestamp, 0).String()
	}

	return
}
