package internal

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNumericChoice(t *testing.T) {
	chk := assert.New(t)
	options := []string{"Option 1", "Option 2", "Option 3"}
	input := "2\n"
	expectedChoice := "Option 2"

	reader := bufio.NewReader(strings.NewReader(input))
	result := GetNumericChoice(reader, options)

	chk.Equal(expectedChoice, result)
}

func TestGetStringInput(t *testing.T) {
	chk := assert.New(t)

	input := "Some input\n"
	expected := "Some input"

	reader := bufio.NewReader(strings.NewReader(input))
	result := GetStringInput(reader, "Enter your name", true)

	chk.Equal(expected, result)
}

func TestParseNumericChoice(t *testing.T) {
	chk := assert.New(t)

	tests := []struct {
		input       string
		max         int
		expectError bool
		expected    int
	}{
		{"1", 3, false, 1},
		{"3", 3, false, 3},
		{"0", 3, true, 0},
		{"4", 3, true, 0},
		{"abc", 3, true, 0},
	}

	for _, test := range tests {
		result, err := parseNumericChoice(test.input, test.max)
		if test.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			chk.Equal(test.expected, result)
		}
	}
}

func TestGetRequirementString(t *testing.T) {
	chk := assert.New(t)

	tests := []struct {
		required bool
		expected string
	}{
		{true, "required"},
		{false, "optional"},
	}

	for _, test := range tests {
		result := getRequirementString(test.required)
		chk.Equal(test.expected, result)
	}
}
