package cmd

import (
	"bufio"
	"fmt"
	"fnc/internal"
	"os"
	"strings"
)

type BranchCommand struct {
	ctx AppContext
}

func NewBranchCommand(ctx AppContext) *BranchCommand {
	return &BranchCommand{ctx: ctx}
}

func (cmd *BranchCommand) Execute(args []string) error {
	reader := bufio.NewReader(os.Stdin)

	internal.ShowAvailableNumericChoices(reader, "Select task type", internal.TaskTypes)

	taskType := internal.GetNumericChoice(reader, internal.TaskTypes)
	taskID := getTaskID(reader)
	description := getDescription(reader)
	prefix := cmd.ctx.config.Prefix
	newBranchName := getNewBranchName(taskType, prefix, taskID, description)
	defaultSourceBranch := internal.GetSourceBranchName(cmd.ctx.config.DefaultBranch)
	internal.CreateNewBranch(newBranchName, defaultSourceBranch)

	return nil
}

func (cmd *BranchCommand) Name() string {
	return "branch"
}

func (cmd *BranchCommand) Description() string {
	return "Create a new branch"
}

func getDescription(reader *bufio.Reader) string {
	out := internal.GetStringInput(reader, "Enter description", true)
	return strings.TrimSpace(strings.ReplaceAll(out, " ", "_"))
}

func getTaskID(reader *bufio.Reader) string {
	out := internal.GetStringInput(reader, "Enter task ID", false)
	return strings.TrimSpace(out)
}

func getNewBranchName(taskType string, prefix string, taskID string, description string) string {
	if taskID == "" || prefix == "" {
		return fmt.Sprintf("%s/%s", taskType, description)
	}
	return fmt.Sprintf("%s/%s-%s_%s", taskType, prefix, taskID, description)
}
