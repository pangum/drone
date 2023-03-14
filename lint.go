package main

type lint struct {
	// 控制程序
	Binary string `default:"${LINT_BINARY=golangci-lint}" json:"binary"`
	// 是否启用
	Enabled *bool `default:"true" json:"enabled"`
	// 超时时间
	Timeout string `default:"10m" json:"timeout"`
	// 启用的Linter
	Linters []string `json:"linters"`
}
