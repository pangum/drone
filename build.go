package main

import (
	"github.com/goexl/gox/field"
)

func (p *plugin) build() (undo bool, err error) {
	for _, _output := range p.Outputs {
		if be := _output.build(p); nil != be {
			err = be
			p.Warn("编译出错", field.New("output", _output))
		}

		if nil != err {
			continue
		}
	}

	return
}
