package main

import (
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/goexl/gfx`
)

func (p *plugin) tidy() (undo bool, err error) {
	_, exists := gfx.Exists(filepath.Join(p.Source, `go.mod`))
	if undo = !exists; undo {
		return
	}

	// 执行清理依赖命令
	err = p.Exec(goExe, drone.Args(`mod`, `tidy`), drone.Dir(p.Source))

	return
}
