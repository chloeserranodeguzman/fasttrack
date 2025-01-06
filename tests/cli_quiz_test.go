package tests

import (
	"bytes"
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

	err := rootCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, output.String(), "Enter your answer (A, B, C, D):")
	assert.Contains(t, output.String(), "You selected: A")
}
