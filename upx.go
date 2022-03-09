package main

import (
	`fmt`
	`strconv`

	`github.com/dronestock/drone`
)

func (p *plugin) upx() (err error) {
	args := []interface{}{
		`--mono`,
		`--color`,
		`-f`,
	}
	if p.Verbose {
		args = append(args, `-v`)
	}

	// 压缩等级
	if _, convErr := strconv.Atoi(p.Compress.Level); nil != convErr {
		args = append(args, fmt.Sprintf(`--%s`, p.Compress.Level))
	} else {
		args = append(args, fmt.Sprintf(`-%s`, p.Compress.Level))
	}
	// 添加输出文件
	args = append(args, p.Output)

	// 执行清理依赖命令
	err = p.Exec(upxExe, drone.Args(args...), drone.Dir(p.Src))

	return
}
