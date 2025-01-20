package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowAvailableNumericChoices(reader *bufio.Reader, description string, options []string) {
	fmt.Printf("%s:\n", description)
	for i, t := range options {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func GetNumericChoice(reader *bufio.Reader, options []string) string {
	for {
		fmt.Print("Enter choice (1-", len(options), "): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		index, err := parseNumericChoice(input, len(options))
		if err != nil {
			fmt.Println("Invalid choice. Try again.")
			continue
		}
		return options[index-1]
	}
}

func GetStringInput(reader *bufio.Reader, inputDescription string, required bool) string {
	for {
		fmt.Printf("%s (%s): ", inputDescription, getRequirementString(required))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && required {
			fmt.Println("The input is required. Try again.")
			continue
		}

		return input
	}
}

func GetMultilineInput(inputDescription string, required bool) string {
	fmt.Printf("%s (%s)\n", inputDescription, getRequirementString(required))
	fmt.Println("Press Ctrl+D (Linux/Mac) or Ctrl+Z (Windows) to finish.")

	var inputLines []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // Stop on EOF (Ctrl+D or Ctrl+Z)
		}
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	return strings.Join(inputLines, "\n")
}

func parseNumericChoice(input string, max int) (int, error) {
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > max {
		return 0, fmt.Errorf("invalid choice")
	}
	return index, nil
}

func getRequirementString(required bool) string {
	if required {
		return "required"
	}
	return "optional"
}
