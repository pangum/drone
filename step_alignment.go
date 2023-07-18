package main

import (
	"context"
	"sync"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

type stepAlignment struct {
	*plugin
}

func newAlignmentStep(plugin *plugin) *stepAlignment {
	return &stepAlignment{
		plugin: plugin,
	}
}

func (a *stepAlignment) Runnable() bool {
	return true
}

func (a *stepAlignment) Run(ctx context.Context) (err error) {
	filenames := make([]string, 0, 1)
	if filenames, err = gfx.All(a.Source, gfx.Suffix(goFileSuffix)); nil != err {
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go a.run(ctx, wg, filename)
	}
	wg.Wait()

	return
}

func (a *stepAlignment) run(_ context.Context, wg *sync.WaitGroup, filename string) {
	defer wg.Done()

	command := a.Command(a.Binary.Alignment)
	command.Args(args.New().Long(strike).Build().Option("fix", filename).Build())
	command.Dir(a.Source)
	environment := command.Environment()
	environment.String(a.envs()...)
	command = environment.Build()
	if _, ee := command.Build().Exec(); nil != ee {
		a.Warn("内存对齐出错", field.New("filename", filename), field.Error(ee))
	} else {
		a.Debug("内存对齐完成", field.New("filename", filename))
	}
}
