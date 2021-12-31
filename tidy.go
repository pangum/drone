package main

import (
	`os`
	`os/exec`
	`path/filepath`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func tidy(conf *config, logger simaqian.Logger) (err error) {
	if exist := gox.IsFileExist(filepath.Join(conf.Input, `go.mod`)); !exist {
		return
	}

	commands := []string{
		`mod`,
		`tidy`,
	}

	// 执行命令
	cmd := exec.Command(`go`, commands...)
	if cmd.Dir, err = filepath.Abs(conf.Input); nil != err {
		return
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, conf.Envs...)
	if conf.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err = cmd.Run(); nil != err {
		logger.Error(`清理依赖出错`, conf.Fields().Connect(field.Error(err))...)
	}

	return
}
