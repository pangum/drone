package main

import (
	"context"
	"path/filepath"

	"github.com/goexl/gfx"
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
	return t.Command(goExe).Args("mod", "tidy").Dir(t.Source).StringEnvs(t.envs()...).Exec()
}
