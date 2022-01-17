package main

import (
	`strings`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func build(conf *config, logger simaqian.Logger) (err error) {
	args := []string{
		`build`,
		`-o`,
		conf.Output,
	}
	if conf.Verbose {
		args = append(args, `-x`)
	}

	// 写入编译标签
	args = append(args, `-ldflags`, strings.Join(conf.flags(), ` `))

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, conf.goExe),
		field.String(`output`, conf.Output),
	}
	logger.Info(`开始编译代码`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(conf.Input))
	if !conf.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(conf.goExe, options...); nil != err {
		logger.Error(`代码编译出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`开始编译成功`, fields...)
	}

	return
}
