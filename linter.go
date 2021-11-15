package main

import (
	`os`
	`os/exec`
	`path/filepath`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func linter(conf *config, logger simaqian.Logger) (err error) {
	commands := []string{
		`run`,
		`--timeout`,
		`10m`,
	}
	for _, linter := range conf.Linters {
		commands = append(commands, `-E`, linter)
	}

	// 执行命令
	cmd := exec.Command(`golangci-lint`, commands...)
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
			`代码检查失败`,
			field.String(`output`, string(output)),
			field.Strings(`command`, commands...),
			field.Error(err),
		)
	} else {
		logger.Info(`代码检查成功`, conf.Fields().Connect(field.Strings(`command`, commands...))...)
	}

	return
}
