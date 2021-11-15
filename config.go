package main

type config struct {
	// 输入文件
	Input string `default:""`
	// 输出文件
	Output string `default:"${}"`
	// 是否启用Lint插件
	Lint bool `default:"true"`
}
