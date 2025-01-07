package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnSelectedAWhenUserSelectsA(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	userInput := bytes.NewBufferString("A\n")

	output := bytes.NewBufferString("")
	rootCmd.SetOut(output)
	rootCmd.SetIn(userInput)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	assert.Contains(t, output.String(), "You selected: A")
}

func TestShouldReturnInvalidInputAndPromptToTryAgain(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	userInput := bytes.NewBufferString("X\n")
	output := bytes.NewBufferString("")

	rootCmd.SetOut(output)
	rootCmd.SetIn(userInput)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	assert.Contains(t, output.String(), "Invalid input. Please enter A, B, C, or D.")

	promptCount := strings.Count(output.String(), "Enter your answer (A, B, C, D):")

	assert.Equal(t, 2, promptCount, "Prompt should appear twice")
}
