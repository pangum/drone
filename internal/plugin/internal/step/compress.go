package step

import (
	"context"

	"github.com/pangum/drone/internal/config"
	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/pangum/drone/internal/core"
)

type Compress struct {
	*internal.Core

	config  *config.Compress
	outputs []*config.Output
	envs    []string
}

func NewCompress(core *internal.Core, config *config.Compress, outputs []*config.Output, envs []string) *Compress {
	return &Compress{
		Core: core,

		config:  config,
		outputs: outputs,
		envs:    envs,
	}
}

func (c *Compress) Runnable() bool {
	return nil != c.config.Enabled && *c.config.Enabled
}

func (c *Compress) Run(_ context.Context) (err error) {
	for _, output := range c.outputs {
		switch c.config.Type {
		case core.CompressTypeUpx:
			err = c.config.Do(&c.Core.Base, &c.Binary, c.Source, c.Dir, c.Verbose, output, c.envs)
		}

		if nil != err {
			continue
		}
	}

	return
}
