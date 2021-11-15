package main

import (
	`strconv`
	`strings`
	`time`

	`github.com/storezhang/god`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	// 输入文件
	Input string `default:"."`
	// 输出文件
	Output string `default:"${DRONE_STAGE_NAME}"`
	// 环境变量
	Envs []string `default:"[\"CGO_ENABLED=0\",\"GOOS=linux\"]"`
	// 是否启用默认配置
	Defaults bool `default:"true"`

	// 是否启用Lint插件
	Lint bool `default:"true"`
	// 启用的Linter
	Linters []string `default:"[\"goerr113\",\"nlreturn\",\"bodyclose\",\"rowserrcheck\",\"gosec\",\"unconvert\",\"misspell\",\"lll\"]"`

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
	c.Input = env(`INPUT`)
	c.Output = env(`OUTPUT`)
	defaults := env(`DEFAULTS`)
	if `` != defaults {
		if c.Defaults, err = strconv.ParseBool(defaults); nil != err {
			return
		}
	}

	lint := env(`LINT`)
	if `` != lint {
		if c.Lint, err = strconv.ParseBool(lint); nil != err {
			return
		}
	}

	c.Name = env(`NAME`)
	c.Version = env(`VERSION`)
	c.Build = env(`BUILD`)
	c.Timestamp = env(`TIMESTAMP`)
	c.Revision = env(`REVISION`)
	c.Branch = env(`BRANCH`)
	if err = god.Set(c); nil != err {
		return
	}

	// 启用默认值
	if !c.Defaults {
		c.Envs = []string{}
		c.Linters = []string{}
	}
	for _, env := range strings.Split(env(`ENVS`), `,`) {
		if `` != env {
			c.Envs = append(c.Envs, env)
		}
	}
	for _, linter := range strings.Split(env(`LINTERS`), `,`) {
		if `` != linter {
			c.Linters = append(c.Linters, linter)
		}
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
