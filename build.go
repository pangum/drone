package main

import (
	`fmt`
	`os`
	`os/exec`
	`path/filepath`
	`strings`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func build(conf *config, logger simaqian.Logger) (err error) {
	commands := []string{
		`build`,
		`-o`,
		conf.Output,
	}
	if conf.Verbose {
		commands = append(commands, `-x`)
	}

	// 写入版本信息
	var ldflags strings.Builder
	ldflags.WriteString(`-s`)
	if `` != conf.Name {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Name=%s'`, conf.Name))
	}
	if `` != conf.Version {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Version=%s'`, conf.Version))
	}
	if `` != conf.Build {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Build=%s'`, conf.Build))
	}
	if `` != conf.Timestamp {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Timestamp=%s'`, conf.Timestamp))
	}
	if `` != conf.Revision {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Revision=%s'`, conf.Revision))
	}
	if `` != conf.Branch {
		ldflags.WriteString(fmt.Sprintf(` -X 'github.com/pangum/pangu.Branch=%s'`, conf.Branch))
	}
	commands = append(commands, `-ldflags`, ldflags.String())

	// 执行命令
	cmd := exec.Command(`go`, commands...)
	if cmd.Dir, err = filepath.Abs(conf.Input); nil != err {
		return
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, conf.Envs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); nil != err {
		logger.Error(`代码编译出错`, conf.Fields().Connect(field.Error(err))...)
	}

	return
}
