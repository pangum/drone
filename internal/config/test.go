package config

type Test struct {
	// 是否启用测试
	Enabled *bool `default:"true" json:"enabled"`
	// 参数
	Args []string `json:"args"`
	// 标志
	Flags []string `json:"flags"`
}
