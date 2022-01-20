package main

import (
	`strings`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) build(logger simaqian.Logger) (undo bool, err error) {
	args := []string{
		`build`,
		`-o`,
		p.config.Output,
	}
	if p.config.Verbose {
		args = append(args, `-x`)
	}

	// 写入编译标签
	args = append(args, `-ldflags`, strings.Join(p.config.flags(), ` `))

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, goExe),
		field.String(`output`, p.config.Output),
	}
	logger.Info(`开始编译代码`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(p.config.Input))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(goExe, options...); nil != err {
		logger.Error(`代码编译出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`开始编译成功`, fields...)
	}

	return
}
