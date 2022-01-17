package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func lint(conf *config, logger simaqian.Logger) (err error) {
	if !conf.Lint {
		return
	}

	args := []string{
		`run`,
		`--timeout`,
		`10m`,
		`--color`,
		`always`,
	}
	// 显示详细信息
	if conf.Verbose {
		args = append(args, `--verbose`)
	}
	// 显示调试信息
	for _, linter := range conf.linters() {
		args = append(args, `--enable`, linter)
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, conf.lintExe),
		field.Strings(`linters`, conf.linters()...),
	}
	logger.Info(`开始代码检查`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(conf.Input))
	if !conf.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(conf.lintExe, options...); nil != err {
		logger.Error(`代码检查出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`代码检查成功`, fields...)
	}

	return
}
