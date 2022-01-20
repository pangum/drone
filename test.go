package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) test(logger simaqian.Logger) (undo bool, err error) {
	if undo = !p.config.Test; undo {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`exe`, goExe),
		field.String(`input`, p.config.Input),
	}
	logger.Info(`开始测试`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(`test`, p.config.Input), gex.Dir(p.config.Input))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(goExe, options...); nil != err {
		logger.Error(`测试出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`测试完成`, fields...)
	}

	return
}
