package main

import (
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func log(conf *config, logger simaqian.Logger, err error) {
	if nil != err {
		logger.Fatal(`编译盘古应用失败`, conf.Fields().Connect(field.Error(err))...)
	} else {
		logger.Info(`编译盘古应用成功`, conf.Fields()...)
	}
}
