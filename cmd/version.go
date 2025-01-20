package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type VersionType struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type VersionCommand struct {
	ctx AppContext
}

func NewVersionCommand(ctx AppContext) *VersionCommand {
	return &VersionCommand{ctx: ctx}
}

func (cmd *VersionCommand) Execute(args []string) error {
	metadata, err := cmd.ctx.metadata.ReadFile("metadata.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var versionStruct struct {
		Version VersionType `json:"version"`
	}

	err = json.Unmarshal(metadata, &versionStruct)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("fnc version %d.%d.%d\n", versionStruct.Version.Major, versionStruct.Version.Minor, versionStruct.Version.Patch)
	return nil
}

func (v *VersionCommand) Name() string {
	return "version"
}

func (v *VersionCommand) Description() string {
	return "Show the application version"
}
