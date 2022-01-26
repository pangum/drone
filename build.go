package main

import (
	`strings`

	`github.com/dronestock/drone`
)

func (p *plugin) build() (undo bool, err error) {
	args := []string{
		`build`,
		`-o`,
		p.Output,
	}
	if p.Verbose {
		args = append(args, `-x`)
	}

	// 写入编译标签
	args = append(args, `-ldflags`, strings.Join(p.flags(), ` `))

	// 执行编译命令
	err = p.Exec(goExe, drone.Args(args...), drone.Dir(p.Input))

	return
}
