package config

type Compress struct {
	// 启用压缩
	Enabled *bool `default:"true" json:"enabled,omitempty"`
	// 压缩等级
	Level string `default:"lzma" json:"level,omitempty" validate:"oneof=1 2 3 4 5 6 7 8 9 best lzma brute ultra-brute"`
}
