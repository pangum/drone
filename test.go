package main

import (
	`github.com/dronestock/drone`
)

type test struct {
	// 是否启用测试
	Enabled bool `default:"true"`
	// 参数
	Args []string
	// 标志
	Flags []string
}

func (p *plugin) test() (undo bool, err error) {
	if undo = !p.Test.Enabled; undo {
		return
	}

	args := []interface{}{
		`test`,
	}
	// 加入默认测试参数
	args = append(args, p.testFlags()...)
	// 加入测试文件
	args = append(args, p.Src)
	// 执行测试命令
	err = p.Exec(goExe, drone.Args(args...), drone.Dir(p.Src))

	return
}
