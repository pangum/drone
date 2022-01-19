package main

import (
	`fmt`
	`strconv`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) upx(logger simaqian.Logger) (undo bool, err error) {
	args := []string{
		`--mono`,
		`--color`,
		`-f`,
	}
	if p.config.Verbose {
		args = append(args, `-v`)
	}

	// 压缩等级
	if _, convErr := strconv.Atoi(p.config.UpxLevel); nil != convErr {
		args = append(args, fmt.Sprintf(`--%s`, p.config.UpxLevel))
	} else {
		args = append(args, fmt.Sprintf(`-%s`, p.config.UpxLevel))
	}

	args = append(args, p.config.Output)

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, p.upxExe),
		field.String(`output`, p.config.Output),
		field.Strings(`args`, args...),
	}
	logger.Info(`开始压缩程序`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(args...), gex.Dir(p.config.Input))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(p.upxExe, options...); nil != err {
		logger.Error(`压缩程序出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`压缩程序成功`, fields...)
	}

	return
}
