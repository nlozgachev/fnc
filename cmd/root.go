package cmd

import (
	"embed"
	"fnc/internal"
	"os"

	"github.com/jessevdk/go-flags"
)

func Run(metadata embed.FS, config internal.Config) int {
	parser := flags.NewNamedParser(internal.AppName, flags.PassDoubleDash|flags.IgnoreUnknown)
	ctx := AppContext{metadata: metadata, config: config}

	commands := registerCommands(ctx, func() { parser.WriteHelp(os.Stdout) })

	for _, cmd := range commands {
		parser.AddCommand(
			cmd.Name(),
			cmd.Description(),
			cmd.Description(),
			cmd,
		)
	}

	/* Inject help command before parsing args */
	if len(os.Args) > 1 && os.Args[1] == "help" {
		parser.WriteHelp(os.Stdout)
		return 0
	}

	_, err := parser.Parse()

	if err != nil {
		parser.WriteHelp(os.Stdout)
		return 1
	}
	return 0
}

func registerCommands(ctx AppContext, writeHelp func()) []Commander {
	return []Commander{
		NewBranchCommand(ctx),
		NewCommitCommand(ctx),
		NewVersionCommand(ctx),
		NewHelpCommand(writeHelp),
	}
}
