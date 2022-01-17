package main

import (
	`path/filepath`

	`github.com/storezhang/gex`
	`github.com/storezhang/gfx`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func tidy(conf *config, logger simaqian.Logger) (err error) {
	mod := filepath.Join(conf.Input, `go.mod`)
	if exist := gfx.Exist(mod); !exist {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`go.mod`, mod),
		field.String(`input`, conf.Input),
	}
	logger.Info(`开户清理依赖`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(`mod`, `tidy`), gex.Dir(conf.Input))
	if !conf.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(conf.goExe, options...); nil != err {
		logger.Error(`清理依赖出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`清理依赖成功`, fields...)
	}

	return
}
