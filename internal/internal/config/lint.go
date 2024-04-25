package config

type Lint struct {
	// 是否启用
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 超时时间
	Timeout string `default:"10m" json:"timeout,omitempty"`
	// 启用的Linter
	Linters []string `json:"linters,omitempty"`
}
