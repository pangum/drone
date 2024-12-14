package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/goexl/args"
	"github.com/pangum/drone/internal/internal/command"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/constant"
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
	if _, se := os.Stat(filepath.Join(t.project.Source, constant.GoModFilename)); nil == se {
		runnable = true
	}

	return
}

func (t *Tidy) Run(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("mod", "tidy").Build()
	err = t.golang.Exec(ctx, arguments)

	return
}
