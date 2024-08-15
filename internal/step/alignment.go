package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
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
	command := a.base.Command(a.binary.Alignment).Context(*ctx)
	command.Args(args.New().Long(constant.Strike).Build().Flag("apply").Subcommand("./...").Build())
	command.Dir(a.project.Source)

	fields := gox.Fields[any]{
		field.New("dir", a.project.Source),
		field.New("binary", a.binary.Alignment),
	}
	if _, ae := command.Build().Exec(); nil != ae {
		err = ae
		a.base.Info("内存对齐出错", fields.Add(field.Error(ae))...)
	} else {
		a.base.Debug("内存对齐完成", fields...)
	}

	return
}
