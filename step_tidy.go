package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
)

type stepTidy struct {
	*plugin
}

func newTidyStep(plugin *plugin) *stepTidy {
	return &stepTidy{
		plugin: plugin,
	}
}

func (t *stepTidy) Runnable() bool {
	_, exists := gfx.Exists(filepath.Join(t.Source, goModFilename))

	return exists
}

func (t *stepTidy) Run(_ context.Context) (err error) {
	command := t.Command(t.Binary.Go)
	command.Args(args.New().Build().Subcommand("mod", "tidy").Build())
	command.Dir(t.Source)
	environment := command.Environment()
	environment.String(t.envs()...)
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}
