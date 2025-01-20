package cmd

type HelpCommand struct {
	writeHelp func()
}

func NewHelpCommand(writeHelp func()) *HelpCommand {
	return &HelpCommand{writeHelp: func() {
		writeHelp()
	}}
}

func (cmd *HelpCommand) Execute(args []string) error {
	args = append(args, "-h")
	return nil
}

func (v *HelpCommand) Name() string {
	return "help"
}

func (v *HelpCommand) Description() string {
	return "Show help"
}
