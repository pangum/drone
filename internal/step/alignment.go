package step

import (
	"context"
	"sync"

	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/core"
	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

type Alignment struct {
	*internal.Core

	config *config.Alignment
	envs   []string
}

func NewAlignment(core *internal.Core, config *config.Alignment, envs []string) *Alignment {
	return &Alignment{
		Core: core,

		config: config,
		envs:   envs,
	}
}

func (a *Alignment) Runnable() bool {
	return nil != a.config.Enabled && *a.config.Enabled
}

func (a *Alignment) Run(ctx *context.Context) (err error) {
	if filenames, ae := gfx.All(a.Source, gfx.Pattern(a.config.Pattern)); nil != ae {
		err = ae
	} else {
		a.run(ctx, filenames)
	}

	return
}

func (a *Alignment) run(ctx *context.Context, filenames []string) {
	wg := new(sync.WaitGroup)
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go a.fix(ctx, wg, filename)
	}
	wg.Wait()
}

func (a *Alignment) fix(_ *context.Context, wg *sync.WaitGroup, filename string) {
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
