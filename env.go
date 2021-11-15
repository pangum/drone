package main

import (
	`fmt`
	`os`
	`strings`
)

// 兼容Drone插件和普通使用
// 优先使用普通模式
// 没有配置再加载Drone配置
func env(envs ...string) (config string) {
	for _, _env := range envs {
		_env = strings.ToUpper(_env)
		if config = eval(_env); `` != config {
			return
		}
	}

	return
}

func eval(config string) (final string) {
	defer func() {
		if final = os.ExpandEnv(final); final == config {
			final = ``
		}
	}()

	var exist bool
	if final, exist = os.LookupEnv(fmt.Sprintf(`PLUGIN_%s`, config)); exist {
		return
	}
	final = config

	return
}
