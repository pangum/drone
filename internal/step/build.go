package step

import (
	"context"
	"sync"

	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gox/field"
)

type Build struct {
	*internal.Core

	outputs []*config.Output
	flags   internal.Flag
	envs    []string
}

func NewBuild(core *internal.Core, outputs []*config.Output, flags internal.Flag, envs []string) *Build {
	return &Build{
		Core: core,

		outputs: outputs,
		flags:   flags,
		envs:    envs,
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx *context.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(b.outputs))
	for _, out := range b.outputs {
		go b.run(ctx, out, wg, &err)
	}
	wg.Wait()

	return
}

func (b *Build) run(_ *context.Context, output *config.Output, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if be := output.Build(&b.Core.Base, &b.Binary, b.Source, b.Dir, b.flags(output.Mode), b.envs); nil != be {
		*err = be
		b.Warn("编译出错", field.New("output", output))
	}
}
