package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/goexl/guc"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/constant"
)

type Alignment struct {
	base   *drone.Base
	binary *config.Binary

	alignment *config.Alignment
	project   *config.Project
}

func NewAlignment(
	base *drone.Base, binary *config.Binary,
	alignment *config.Alignment, project *config.Project,
) *Alignment {
	return &Alignment{
		base:   base,
		binary: binary,

		alignment: alignment,
		project:   project,
	}
}

func (a *Alignment) Runnable() bool {
	return nil != a.alignment.Enabled && *a.alignment.Enabled
}

func (a *Alignment) Run(ctx *context.Context) (err error) {
	if filenames, ae := gfx.All(a.project.Source, gfx.Pattern(a.alignment.Pattern)); nil != ae {
		err = ae
	} else {
		a.run(ctx, filenames, &err)
	}

	return
}

func (a *Alignment) run(ctx *context.Context, filenames []string, err *error) {
	wg := new(guc.WaitGroup)
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go a.fix(ctx, wg, filename, err)
	}
	wg.Wait()
}

func (a *Alignment) fix(ctx *context.Context, wg *guc.WaitGroup, filename string, err *error) {
	defer wg.Done()

	command := a.base.Command(a.binary.Alignment).Context(*ctx)
	command.Args(args.New().Long(constant.Strike).Build().Option("fix", filename).Build())
	command.Dir(a.project.Source)
	if _, ae := command.Build().Exec(); nil != ae {
		*err = ae
		a.base.Info("内存对齐出错", field.New("filename", filename))
	} else {
		a.base.Debug("内存对齐完成", field.New("filename", filename))
	}
}
