package step

import (
	"context"
	"github.com/dronestock/drone"
	"github.com/pangum/drone/internal/plugin/internal"
	"sync"

	"github.com/goexl/gox/field"
	"github.com/pangum/drone/internal/config"
)

type Build struct {
	drone.Base
	internal.Core

	outputs []*config.Output
	flags   internal.Flag
	envs    []string
}

func NewBuild(base drone.Base, core internal.Core, flags internal.Flag, envs []string) *Build {
	return &Build{
		Base: base,
		Core: core,

		flags: flags,
		envs:  envs,
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx context.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(b.outputs))
	for _, out := range b.outputs {
		go b.build(ctx, out, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (b *Build) build(_ context.Context, output *config.Output, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	if be := output.Build(&b.Base, &b.Binary, b.Source, b.Dir, b.flags(output.Mode), b.envs); nil != be {
		*err = be
		b.Warn("编译出错", field.New("output", output))
	}
}
