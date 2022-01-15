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
		`--color`,
		`always`,
	}
	if conf.Verbose {
		commands = append(commands, `--verbose`)
	}
	for _, _linter := range conf.Linters {
		commands = append(commands, `--enable`, _linter)
	}

	// 执行命令
	cmd := exec.Command(`golangci-lint`, commands...)
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
		logger.Error(`代码检查出错`, conf.Fields().Connect(field.Error(err))...)
	}

	return
}
