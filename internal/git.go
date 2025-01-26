package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CreateNewBranch(newBranch string, sourceBranch string) {
	cmd := exec.Command("git", "checkout", "-b", newBranch, sourceBranch)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error creating the new branch:", err)
		return
	}

	fmt.Printf("Branch \"%s\" created from \"%s\"\n", newBranch, sourceBranch)
}

func CreateCommit(fullCommitMessage string, commitBody string, noVerify bool) {
	commitCmd := exec.Command("git", "commit", "-m", fullCommitMessage)
	if commitBody != "" {
		commitCmd.Args = append(commitCmd.Args, "-m", commitBody)
	}

	if noVerify {
		commitCmd.Args = append(commitCmd.Args, "--no-verify")
	}

	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr

	err := commitCmd.Run()
	if err != nil {
		fmt.Println("Error executing git commit:", err)
		return
	}

	fmt.Println("Commit successful.")
}

func GetCurrentBranchName() string {
	branchNameCmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	branchNameBytes, err := branchNameCmd.Output()

	if err != nil {
		fmt.Println("Error getting current branch name:", err)
		return ""
	}

	return strings.TrimSpace(string(branchNameBytes))
}

func GetSourceBranchName(defaultSourceBranch string) string {
	if defaultSourceBranch == "" {
		return GetCurrentBranchName()
	}
	return defaultSourceBranch
}

func ExtractTaskIDFromBranch() string {
	branchName := GetCurrentBranchName()
	return ExtractBranchPrefix(branchName)
}
