package internal

import "github.com/pangum/drone/internal/core"

type Flag func(mode core.Mode) []string
