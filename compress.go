package main

type compress struct {
	// 启用压缩
	Enabled *bool `default:"true" json:"enabled"`
	// 类型
	Type string `default:"upx" json:"type" validate:"oneof=upx"`
	// 压缩等级
	Level string `default:"lzma" json:"level" validate:"oneof=1 2 3 4 5 6 7 8 9 best lzma brute ultra-brute"`
}

func (p *plugin) compress() (undo bool, err error) {
	if undo = !*p.Compress.Enabled; undo {
		return
	}

	switch p.Compress.Type {
	case compressTypeUpx:
		err = p.upx()
	}

	return
}
