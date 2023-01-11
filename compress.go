package main

type compress struct {
	// 启用压缩
	Enabled *bool `default:"true" json:"enabled"`
	// 类型
	Type compressType `default:"upx" json:"type" validate:"oneof=upx"`
	// 压缩等级
	Level string `default:"lzma" json:"level" validate:"oneof=1 2 3 4 5 6 7 8 9 best lzma brute ultra-brute"`
}

func (p *plugin) compress() (undo bool, err error) {
	if undo = !*p.Compress.Enabled; undo {
		return
	}

	for _, _output := range p.Outputs {
		switch p.Compress.Type {
		case compressTypeUpx:
			err = p.Compress.upx(p, _output)
		}

		if nil != err {
			continue
		}
	}

	return
}
