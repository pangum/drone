package main

import (
	`bytes`
	`fmt`
	`os`
	`os/exec`
	`path/filepath`

	`github.com/storezhang/gox`
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
