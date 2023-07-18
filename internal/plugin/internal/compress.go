package internal

import (
	"context"

	"github.com/pangum/drone/internal/core"
	"github.com/pangum/drone/internal/plugin"
)

type Compress struct {
	*plugin.Plugin
}

func NewCompress(plugin *plugin.Plugin) *Compress {
	return &Compress{
		Plugin: plugin,
	}
}

func (c *Compress) Runnable() bool {
	return c.CompressEnabled()
}

func (c *Compress) Run(_ context.Context) (err error) {
	for _, output := range c.Outputs {
		switch c.Compress.Type {
		case core.CompressTypeUpx:
			err = c.Compress.Do(&c.Base, &c.Binary, c.Source, c.Dir, c.Verbose, output, c.Environments())
		}

		if nil != err {
			continue
		}
	}

	return
}
