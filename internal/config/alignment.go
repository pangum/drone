package config

type Alignment struct {
	// 是否启用
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 需要对齐的文件
	Pattern string `default:"*.go" json:"pattern,omitempty"`
}
