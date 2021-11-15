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
	for _, env := range conf.Envs {
		cmd.Env = append(cmd.Env, env)
	}
	if err = cmd.Run(); nil != err {
		output, _ := cmd.CombinedOutput()
		logger.Warn(
			`清理依赖出错`,
			field.String(`output`, string(output)),
			field.Strings(`command`, commands...),
			field.Error(err),
		)
	} else {
		logger.Info(`清理依赖成功`, conf.Fields().Connect(field.Strings(`command`, commands...))...)
	}

	return
}
