package main

import (
	`path/filepath`

	`github.com/storezhang/gex`
	`github.com/storezhang/gfx`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) tidy(logger simaqian.Logger) (undo bool, err error) {
	mod := filepath.Join(p.config.Input, `go.mod`)
	if undo = !gfx.Exist(mod); undo {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`go.mod`, mod),
		field.String(`input`, p.config.Input),
	}
	logger.Info(`开始清理依赖`, fields...)

	// 执行命令
	options := gex.NewOptions(gex.Args(`mod`, `tidy`), gex.Dir(p.config.Input))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(p.goExe, options...); nil != err {
		logger.Error(`清理依赖出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`清理依赖成功`, fields...)
	}

	return
}
