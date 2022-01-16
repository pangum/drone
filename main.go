package main

import (
	`github.com/storezhang/simaqian`
)

func main() {
	var err error
	// 有错误，输出错误日志
	var logger simaqian.Logger
	if logger, err = simaqian.New(); nil != err {
		panic(err)
	}

	// 取各种参数
	conf := new(config)
	defer func() {
		log(conf, logger, err)
	}()
	if err = conf.load(); nil != err {
		return
	}

	// 记录配置日志信息
	logger.Info(`加载配置完成`, conf.Fields()...)

	// 更新依赖
	if err = tidy(conf, logger); nil != err {
		return
	}
	// 代码检查
	if conf.Lint {
		if err = lint(conf, logger); nil != err {
			return
		}
	}
	// 编译
	if err = build(conf, logger); nil != err {
		return
	}
	// 压缩
	if err = upx(conf, logger); nil != err {
		return
	}
}
