package main

import (
	`encoding/json`
	`os`
	`strconv`
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/mengpo`
)

var configJson = `{
	"input": "${PLUGIN_INPUT=$INPUT}",
	"output": "${PLUGIN_OUTPUT=$OUTPUT}",
	"envs": [${PLUGIN_ENVS=$ENVS}],
	"defaults": ${PLUGIN_DEFAULTS=$DEFAULTS},
	"verbose": ${PLUGIN_VERBOSE=$VERBOSE},

	"lint": ${PLUGIN_LINT=$LINT},
	"linters": [${PLUGIN_LINTERS=$LINTERS}],

	"name": "${PLUGIN_NAME=$NAME}",
	"version": "${PLUGIN_VERSION=$VERSION}",
	"build": "${PLUGIN_BUILD=$BUILD}",
	"timestamp": "${PLUGIN_TIMESTAMP=$TIMESTAMP}",
	"revision": "${PLUGIN_REVISION=$REVISION}",
	"branch": "${PLUGIN_BRANCH=$BRANCH}"
}
`

type config struct {
	// 输入文件
	Input string `default:"."`
	// 输出文件
	Output string `default:"${DRONE_STAGE_NAME}"`
	// 环境变量
	Envs []string `default:"['CGO_ENABLED=0','GOOS=linux']"`
	// 是否启用默认配置
	Defaults bool `default:"true"`
	// 是否显示调试信息
	Verbose bool `default:"false"`

	// 是否启用Lint插件
	Lint bool `default:"true"`
	// 启用的Linter
	Linters []string `default:"['goerr113','nlreturn','bodyclose','rowserrcheck','gosec','unconvert','misspell','lll']"`

	// 应用名称
	Name string `default:"$DRONE_STAGE_NAME"`
	// 应用版本
	Version string `default:"${DRONE_TAG=${DRONE_COMMIT_BRANCH:latest}"`
	// 编译版本
	Build string `default:"${DRONE_BUILD_NUMBER}"`
	// 编译时间
	Timestamp string `default:"${DRONE_BUILD_STARTED}"`
	// 分支版本
	Revision string `default:"${DRONE_COMMIT_SHA}"`
	// 分支
	Branch string `default:"${DRONE_COMMIT_BRANCH}"`
}

func (c *config) load() (err error) {
	// 处理环境变量
	if err = parseEnvs(`ENVS`, `LINTERS`); nil != err {
		return
	}
	configJson = os.ExpandEnv(configJson)
	if err = json.Unmarshal([]byte(configJson), c); nil != err {
		return
	}
	if err = mengpo.Set(c); nil != err {
		return
	}

	// 启用默认值
	if c.Defaults {
		c.Envs = append(c.Envs, ``)
		c.Linters = []string{}
	}

	// 将时间变换成易读形式
	if timestamp, parseErr := strconv.ParseInt(c.Timestamp, 10, 64); nil == parseErr {
		c.Timestamp = time.Unix(timestamp, 0).String()
	}

	return
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
