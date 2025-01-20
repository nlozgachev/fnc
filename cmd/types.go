package cmd

import (
	"embed"
	"fnc/internal"
)

type AppContext struct {
	metadata embed.FS
	config   internal.Config
}

type Commander interface {
	Execute(args []string) error
	Name() string
	Description() string
}
