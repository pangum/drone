package step

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/pangum/drone/internal/core"
	"github.com/pangum/drone/internal/plugin"
)

type Tidy struct {
	*plugin.Plugin
}

func NewTidy(plugin *plugin.Plugin) *Tidy {
	return &Tidy{
		Plugin: plugin,
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
	environment.String(t.Environments()...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
