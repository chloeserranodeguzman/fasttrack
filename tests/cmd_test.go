package tests

import (
	"bytes"
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestQuizCommand(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	output := bytes.NewBufferString("")
	rootCmd.SetOut(output)
	rootCmd.SetArgs([]string{"quiz"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	expectedOutput := `Question: What is the capital of Japan?
A) Tokyo
B) Kyoto
C) Osaka
D) Nagoya
`

	assert.Equal(t, expectedOutput, output.String())
}
