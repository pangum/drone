package main

import (
	`fmt`
	`strconv`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func upx(conf *config, logger simaqian.Logger) (err error) {
	args := []string{
		`-l`,
		`--mono`,
		`--color`,
	}
	if conf.Verbose {
		args = append(args, `-v`)
	}

	// 压缩等级
	if _, convErr := strconv.Atoi(conf.UpxLevel); nil != convErr {
		args = append(args, fmt.Sprintf(`--%s`, conf.UpxLevel))
	} else {
		args = append(args, fmt.Sprintf(`-%s`, conf.UpxLevel))
	}

	args = append(args, conf.Output)

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, conf.upxExe),
		field.String(`output`, conf.Output),
	}
	logger.Info(`开始压缩程序`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(conf.Input))
	if !conf.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(conf.upxExe, options...); nil != err {
		logger.Error(`压缩程序出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`压缩程序成功`, fields...)
	}

	return
}
