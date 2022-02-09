package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) test() (undo bool, err error) {
	if undo = !p.Test; undo {
		return
	}

	args := []string{
		`test`,
	}
	// 加入默认测试参数
	args = append(args, p.testFlags()...)
	// 加入测试文件
	args = append(args, p.Input)
	// 执行测试命令
	err = p.Exec(goExe, drone.Args(args...), drone.Dir(p.Input))

	return
}
