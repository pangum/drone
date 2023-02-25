package main

import (
	"context"
	"sync"

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

func (b *stepBuild) Run(ctx context.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(b.Outputs))
	for _, out := range b.Outputs {
		go b.build(ctx, out, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (b *stepBuild) build(_ context.Context, output *output, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	if be := output.build(b.plugin); nil != be {
		*err = be
		b.Warn("编译出错", field.New("output", output))
	}
}
