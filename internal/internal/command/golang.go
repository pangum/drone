package command

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/goexl/args"
	"github.com/goexl/gox"
	"github.com/pangum/drone/internal/internal/config"
	"github.com/pangum/drone/internal/internal/constant"
	"github.com/pangum/drone/internal/internal/core"
)

type Golang struct {
	*drone.Base

	binary  *config.Binary
	project *config.Project

	defaultEnvironments []string
}

func NewGolang(base *drone.Base, binary *config.Binary, project *config.Project) *Golang {
	return &Golang{
		Base: base,

		binary:  binary,
		project: project,

		defaultEnvironments: []string{
			"CGO_ENABLED=0",
		},
	}
}

func (g *Golang) Exec(ctx *context.Context, arguments *args.Arguments, environments ...*core.Environment) (err error) {
	if g.Verbose {
		arguments = arguments.Rebuild().Flag("x").Build()
	}

	command := g.Command(g.binary.Go).Context(*ctx)
	command.Args(arguments)
	command.Dir(g.project.Source)
	environment := command.Environment()
	environment.String(g.environments()...)
	for _, env := range environments {
		environment.Kv(env.Key(), env.Value())
	}
	command = environment.Build()
	_, err = command.Build().Exec()

	return
}

func (g *Golang) environments() (environments []string) {
	environments = make([]string, 0, len(g.project.Environments)+2)
	if g.Default() {
		environments = append(environments, g.defaultEnvironments...)
	}
	for _, private := range g.project.Privates {
		goPrivate := gox.StringBuilder(constant.GoPrivate, constant.Equal, private).String()
		goNoProxy := gox.StringBuilder(constant.GoNoProxy, constant.Equal, private).String()
		environments = append(environments, goPrivate, goNoProxy)
	}
	environments = append(environments, g.project.Environments...)

	return
}
