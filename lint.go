package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) lint(logger simaqian.Logger) (undo bool, err error) {
	if undo = !p.config.Lint; undo {
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
	if p.config.Verbose {
		args = append(args, `--verbose`)
	}
	// 显示调试信息
	for _, linter := range p.config.linters() {
		args = append(args, `--enable`, linter)
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, lintExe),
		field.Strings(`linters`, p.config.linters()...),
	}
	logger.Info(`开始代码检查`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(p.config.Input))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(lintExe, options...); nil != err {
		logger.Error(`代码检查出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`代码检查成功`, fields...)
	}

	return
}