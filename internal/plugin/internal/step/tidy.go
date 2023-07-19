package step

import (
	"context"
	"path/filepath"

	"github.com/pangum/drone/internal/plugin/internal"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
)

type Tidy struct {
	*internal.Core

	envs []string
}

func NewTidy(core *internal.Core, envs []string) *Tidy {
	return &Tidy{
		Core: core,

		envs: envs,
	}
}

func (t *Tidy) Runnable() bool {
	_, exists := gfx.Exists(filepath.Join(t.Source, core.GoModFilename))

	return exists
}

func (t *Tidy) Run(_ context.Context) (err error) {
	command := t.Command(t.Binary.Go)
	command.Args(args.New().Build().Subcommand("mod", "tidy").Build())
	command.Dir(t.Source)
	environment := command.Environment()
	environment.String(t.envs...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
