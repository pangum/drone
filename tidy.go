package main

import (
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gfx`
)

func (p *plugin) tidy() (undo bool, err error) {
	if undo = !gfx.Exist(filepath.Join(p.Source, `go.mod`)); undo {
		return
	}

	// 执行清理依赖命令
	err = p.Exec(goExe, drone.Args(`mod`, `tidy`), drone.Dir(p.Source))

	return
}
