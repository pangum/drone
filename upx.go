package main

import (
	`fmt`
	`strconv`

	`github.com/dronestock/drone`
)

func (p *plugin) upx() (undo bool, err error) {
	args := []interface{}{
		`--mono`,
		`--color`,
		`-f`,
	}
	if p.Verbose {
		args = append(args, `-v`)
	}

	// 压缩等级
	if _, convErr := strconv.Atoi(p.UpxLevel); nil != convErr {
		args = append(args, fmt.Sprintf(`--%s`, p.UpxLevel))
	} else {
		args = append(args, fmt.Sprintf(`-%s`, p.UpxLevel))
	}
	// 添加输出文件
	args = append(args, p.Output)

	// 执行清理依赖命令
	err = p.Exec(upxExe, drone.Args(args...), drone.Dir(p.Input))

	return
}
