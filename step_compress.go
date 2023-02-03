package main

import (
	"context"
)

type stepCompress struct {
	*plugin
}

func newCompressStep(plugin *plugin) *stepCompress {
	return &stepCompress{
		plugin: plugin,
	}
}

func (c *stepCompress) Runnable() bool {
	return c.compressEnabled()
}

func (c *stepCompress) Run(_ context.Context) (err error) {
	for _, _output := range c.Outputs {
		switch c.Compress.Type {
		case compressTypeUpx:
			err = c.Compress.upx(c.plugin, _output)
		}

		if nil != err {
			continue
		}
	}

	return
}
