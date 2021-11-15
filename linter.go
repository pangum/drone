package main

import (
	`bytes`
	`fmt`
	`os`
	`os/exec`
	`path/filepath`

	`github.com/storezhang/simaqian`
)

func linter(conf *config, logger simaqian.Logger) (err error) {
	commands := []string{
		`run`,
		`--timeout`,
		`10m`,
		`--verbose`,
		`--color`,
		`always`,
	}
	for _, linter := range conf.Linters {
		commands = append(commands, `--enable`, linter)
	}

	// 执行命令
	cmd := exec.Command(`golangci-lint`, commands...)
	if cmd.Dir, err = filepath.Abs(conf.Input); nil != err {
		return
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, conf.Envs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Run(); nil != err {
		fmt.Println(stderr.String())
	} else {
		fmt.Println(stdout.String())
	}

	return
}
