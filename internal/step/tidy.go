package step

import (
	"context"

	"github.com/goexl/args"
	"github.com/pangum/drone/internal/internal/command"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/constant"

	"github.com/goexl/gfx"
)

type Tidy struct {
	golang  *command.Golang
	project *config.Project
}

func NewTidy(golang *command.Golang, project *config.Project) *Tidy {
	return &Tidy{
		golang:  golang,
		project: project,
	}
}

func (t *Tidy) Runnable() (runnable bool) {
	_, exists := gfx.Exists().Dir(t.project.Source).Filename(constant.GoModFilename).Build().Check()
	runnable = exists

	return
}

func (t *Tidy) Run(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("mod", "tidy").Build()
	err = t.golang.Exec(ctx, arguments)

	return
}
