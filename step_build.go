package main

import (
	"context"

	"github.com/goexl/gox/field"
)

type stepBuild struct {
	*plugin
}

func newBuildStep(plugin *plugin) *stepBuild {
	return &stepBuild{
		plugin: plugin,
	}
}

func (b *stepBuild) Runnable() bool {
	return true
}

func (b *stepBuild) Run(_ context.Context) (err error) {
	for _, _output := range b.Outputs {
		if be := _output.build(b.plugin); nil != be {
			err = be
			b.Warn("编译出错", field.New("output", _output))
		}

		if nil != err {
			continue
		}
	}

	return
}
