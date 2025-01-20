package main

import (
	"embed"
	"fnc/cmd"
	"fnc/internal"
	"os"
)

//go:embed metadata.json
var metadata embed.FS

func main() {
	config, _ := internal.GetConfig()
	os.Exit(cmd.Run(metadata, config))
}
