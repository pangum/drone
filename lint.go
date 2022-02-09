package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) lint() (undo bool, err error) {
	if undo = !p.Lint; undo {
		return
	}

	args := []interface{}{
		`run`,
		`--timeout`,
		`10m`,
		`--color`,
		`always`,
	}
	// 显示详细信息
	if p.Verbose {
		args = append(args, `--verbose`)
	}
	// 显示调试信息
	for _, linter := range p.linters() {
		args = append(args, `--enable`, linter)
	}

	// 执行代码检查命令
	err = p.Exec(lintExe, drone.Args(args...), drone.Dir(p.Input))

	return
}
