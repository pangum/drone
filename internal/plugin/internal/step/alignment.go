package step

import (
	"context"
	"sync"

	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
	"github.com/pangum/drone/internal/core"
)

type Alignment struct {
	*internal.Core

	envs []string
}

func NewAlignment(core *internal.Core, envs []string) *Alignment {
	return &Alignment{
		Core: core,

		envs: envs,
	}
}

func (a *Alignment) Runnable() bool {
	return true
}

func (a *Alignment) Run(ctx context.Context) (err error) {
	filenames := make([]string, 0, 1)
	if filenames, err = gfx.All(a.Source, gfx.Suffix(core.GoFileSuffix)); nil != err {
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

func (a *Alignment) run(_ context.Context, wg *sync.WaitGroup, filename string) {
	defer wg.Done()

	command := a.Command(a.Binary.Alignment)
	command.Args(args.New().Long(core.Strike).Build().Option("fix", filename).Build())
	command.Dir(a.Source)
	environment := command.Environment()
	environment.String(a.envs...)
	command = environment.Build()
	if _, ee := command.Build().Exec(); nil != ee {
		a.Info("内存对齐出错", field.New("filename", filename))
	} else {
		a.Debug("内存对齐完成", field.New("filename", filename))
	}
}
