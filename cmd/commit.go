package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fnc/internal"
)

type CommitCommand struct {
	ctx AppContext
}

func NewCommitCommand(ctx AppContext) *CommitCommand {
	return &CommitCommand{ctx: ctx}
}
func (cmd *CommitCommand) Execute(args []string) error {
	reader := bufio.NewReader(os.Stdin)

	internal.ShowAvailableNumericChoices(reader, "Select the type of change you're committing", internal.CommitTypes)

	commitType := getCommitTypeUserInput(reader)
	commitMessage := getCommitMessageUserInput(reader)
	commitBody := getCommitBodyUserInput(reader)
	taskID := internal.ExtractTaskIDFromBranch()
	fullCommitMessage := assembleCommitMessage(commitType, commitMessage, taskID)

	internal.CreateCommit(fullCommitMessage, commitBody)

	return nil
}

func (c *CommitCommand) Name() string {
	return "commit"
}

func (c *CommitCommand) Description() string {
	return "Create a new commit"
}

func getCommitMessageUserInput(reader *bufio.Reader) string {
	commitMessage := internal.GetStringInput(reader, "Enter the commit message", true)
	commitMessage = strings.TrimSpace(commitMessage)
	return commitMessage
}

func getCommitBodyUserInput(reader *bufio.Reader) string {
	commitBody := internal.GetStringInput(reader, "Enter the commit body", false)
	commitBody = strings.TrimSpace(commitBody)
	return commitBody
}

func getCommitTypeUserInput(reader *bufio.Reader) string {
	commitType := internal.GetNumericChoice(reader, internal.CommitTypes)
	return commitType
}

func assembleCommitMessage(commitType string, commitMessage string, id string) string {
	fullCommitMessage := fmt.Sprintf("%s: %s", commitType, commitMessage)
	if id != "" {
		fullCommitMessage = fmt.Sprintf("%s (%s)", fullCommitMessage, id)
	}

	return fullCommitMessage
}
