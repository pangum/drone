package internal

import (
	"context"
	"sync"

	"github.com/goexl/gox/field"
	"github.com/pangum/drone/internal/config"
	"github.com/pangum/drone/internal/plugin"
)

type Build struct {
	*plugin.Plugin
}

func NewBuild(plugin *plugin.Plugin) *Build {
	return &Build{
		Plugin: plugin,
	}
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Run(ctx context.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(b.Outputs))
	for _, out := range b.Outputs {
		go b.build(ctx, out, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (b *Build) build(_ context.Context, output *config.Output, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	if be := output.Build(&b.Plugin.Base, &b.Binary, b.Source, b.Dir, b.Flags(output.Mode), b.Environments()); nil != be {
		*err = be
		b.Warn("编译出错", field.New("output", output))
	}
}
