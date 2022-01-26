package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) test() (undo bool, err error) {
	if undo = !p.Test; undo {
		return
	}

	// 执行测试命令
	err = p.Exec(goExe, drone.Args(`test`, p.Input), drone.Dir(p.Input))

	return
}
